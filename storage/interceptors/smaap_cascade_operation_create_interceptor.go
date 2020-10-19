/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package interceptors

import (
	"context"
	"github.com/Peripli/service-manager/operations"
	"github.com/Peripli/service-manager/pkg/query"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/storage"
	"time"
)

const CascadeOperationCreateInterceptorProviderName = "CascadeOperationCreateInterceptorProvider"

type cascadeOperationCreateInterceptor struct {
}

type CascadeOperationCreateInterceptorProvider struct {
}

func (c *CascadeOperationCreateInterceptorProvider) Provide() storage.CreateOnTxInterceptor {
	return &cascadeOperationCreateInterceptor{}
}

func (c *CascadeOperationCreateInterceptorProvider) Name() string {
	return CascadeOperationCreateInterceptorProviderName
}

func (co *cascadeOperationCreateInterceptor) OnTxCreate(f storage.InterceptCreateOnTxFunc) storage.InterceptCreateOnTxFunc {
	return func(ctx context.Context, storage storage.Repository, obj types.Object) (types.Object, error) {
		operation := obj.(*types.Operation)
		isVirtual := types.IsVirtualType(operation.ResourceType)
		if isVirtual || operation.CascadeRootID == "" || operation.Type != types.DELETE {
			return f(ctx, storage, operation)
		}

		// init operation properties
		operation.PlatformID = types.SMPlatform
		operation.State = types.PENDING
		operation.Base.CreatedAt = time.Now()
		operation.Base.UpdatedAt = time.Now()
		operation.Base.Ready = true

		if err := operation.Validate(); err != nil {
			return nil, err
		}

		if duplicate, err := doesExistCascadeOperationForResource(ctx, storage, operation); err != nil || duplicate != nil {
			// in case cascade operation does exists for this resource
			return duplicate, err
		}

		cascadeResource, err := storage.Get(ctx, operation.ResourceType, query.ByField(query.EqualsOperator, "id", operation.ResourceID))
		if err != nil {
			return nil, err
		}

		ops, err := operations.GetAllLevelsCascadeOperations(ctx, cascadeResource, operation, storage)
		if err != nil {
			return nil, err
		}

		for _, op := range ops {
			if _, err := storage.Create(ctx, op); err != nil {
				return nil, util.HandleStorageError(err, string(op.GetType()))
			}
		}

		return f(ctx, storage, operation)
	}
}

func doesExistCascadeOperationForResource(ctx context.Context, storage storage.Repository, operation *types.Operation) (*types.Operation, error) {
	// if same resource exists in other cascade tree the existing tree/subtree root will be returned
	// in case of service instance deletion failure errors might be lost but will be available in the original tree
	criteria := []query.Criterion{
		query.ByField(query.EqualsOperator, "resource_id", operation.ResourceID),
		query.ByField(query.InOperator, "state", string(types.IN_PROGRESS), string(types.PENDING)),
		query.ByField(query.NotEqualsOperator, "cascade_root_id", ""),
	}
	op, err := storage.Get(ctx, types.OperationType, criteria...)
	if err != nil {
		if err == util.ErrNotFoundInStorage {
			return nil, nil
		}
		return nil, err
	}
	return op.(*types.Operation), nil
}
