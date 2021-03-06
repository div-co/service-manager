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

package filters

import (
	"net/http"

	"github.com/Peripli/service-manager/pkg/log"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/storage"

	"github.com/Peripli/service-manager/pkg/web"
	"github.com/tidwall/gjson"
)

const planIDProperty = "service_plan_id"

const ServiceInstanceVisibilityFilterName = "ServiceInstanceVisibilityFilter"

// serviceInstanceVisibilityFilter ensures that the tenant making the provisioning/update request
// has the necessary visibilities - i.e. that he has the right to consume the requested plan.
type serviceInstanceVisibilityFilter struct {
	repository                    storage.Repository
	getInstanceVisibilityMetadata func(req *web.Request, repository storage.Repository) (*VisibilityMetadata, error)
}

// NewServiceInstanceVisibilityFilter creates a new serviceInstanceVisibilityFilter filter
func NewServiceInstanceVisibilityFilter(repository storage.Repository, getInstanceVisibilityMetadata func(req *web.Request, repository storage.Repository) (*VisibilityMetadata, error)) *serviceInstanceVisibilityFilter {
	return &serviceInstanceVisibilityFilter{
		repository:                    repository,
		getInstanceVisibilityMetadata: getInstanceVisibilityMetadata,
	}
}

func (*serviceInstanceVisibilityFilter) Name() string {
	return ServiceInstanceVisibilityFilterName
}

func (f *serviceInstanceVisibilityFilter) Run(req *web.Request, next web.Handler) (*web.Response, error) {
	ctx := req.Context()

	visibilityMetadata, err := f.getInstanceVisibilityMetadata(req, f.repository)
	if err != nil {
		return nil, err
	}

	if visibilityMetadata.LabelKey == "" || visibilityMetadata.LabelValue == "" || req.Method == http.MethodDelete {
		return next.Handle(req)
	}

	planID := gjson.GetBytes(req.Body, planIDProperty).String()

	if planID == "" {
		log.C(ctx).Info("Plan ID is not provided in the request. Proceeding with the next handler...")
		return next.Handle(req)
	}

	list, err := f.repository.QueryForList(ctx, types.VisibilityType, storage.QueryForVisibilityWithPlatformAndPlan, map[string]interface{}{
		"platform_id":     visibilityMetadata.PlatformID,
		"service_plan_id": planID,
		"key":             visibilityMetadata.LabelKey,
		"val":             visibilityMetadata.LabelValue,
	})

	if err != nil {
		return nil, util.HandleStorageError(err, types.VisibilityType.String())
	}

	if list.Len() > 0 {
		return next.Handle(req)
	}

	visibilityError := &util.HTTPError{
		ErrorType:   "NotFound",
		Description: "could not find such service plan",
		StatusCode:  http.StatusNotFound,
	}

	return nil, visibilityError
}

func (*serviceInstanceVisibilityFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{
		{
			Matchers: []web.Matcher{
				web.Path(web.ServiceInstancesURL + "/**"),
				web.Methods(http.MethodPost, http.MethodPatch, http.MethodDelete),
			},
		},
	}
}
