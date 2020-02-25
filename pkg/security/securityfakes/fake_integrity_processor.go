// Code generated by counterfeiter. DO NOT EDIT.
package securityfakes

import (
	"sync"

	"github.com/Peripli/service-manager/pkg/security"
)

type FakeIntegrityProcessor struct {
	CalculateIntegrityStub        func(security.IntegralObject) ([]byte, error)
	calculateIntegrityMutex       sync.RWMutex
	calculateIntegrityArgsForCall []struct {
		arg1 security.IntegralObject
	}
	calculateIntegrityReturns struct {
		result1 []byte
		result2 error
	}
	calculateIntegrityReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ValidateIntegrityStub        func(security.IntegralObject) bool
	validateIntegrityMutex       sync.RWMutex
	validateIntegrityArgsForCall []struct {
		arg1 security.IntegralObject
	}
	validateIntegrityReturns struct {
		result1 bool
	}
	validateIntegrityReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIntegrityProcessor) CalculateIntegrity(arg1 security.IntegralObject) ([]byte, error) {
	fake.calculateIntegrityMutex.Lock()
	ret, specificReturn := fake.calculateIntegrityReturnsOnCall[len(fake.calculateIntegrityArgsForCall)]
	fake.calculateIntegrityArgsForCall = append(fake.calculateIntegrityArgsForCall, struct {
		arg1 security.IntegralObject
	}{arg1})
	fake.recordInvocation("CalculateIntegrity", []interface{}{arg1})
	fake.calculateIntegrityMutex.Unlock()
	if fake.CalculateIntegrityStub != nil {
		return fake.CalculateIntegrityStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.calculateIntegrityReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIntegrityProcessor) CalculateIntegrityCallCount() int {
	fake.calculateIntegrityMutex.RLock()
	defer fake.calculateIntegrityMutex.RUnlock()
	return len(fake.calculateIntegrityArgsForCall)
}

func (fake *FakeIntegrityProcessor) CalculateIntegrityCalls(stub func(security.IntegralObject) ([]byte, error)) {
	fake.calculateIntegrityMutex.Lock()
	defer fake.calculateIntegrityMutex.Unlock()
	fake.CalculateIntegrityStub = stub
}

func (fake *FakeIntegrityProcessor) CalculateIntegrityArgsForCall(i int) security.IntegralObject {
	fake.calculateIntegrityMutex.RLock()
	defer fake.calculateIntegrityMutex.RUnlock()
	argsForCall := fake.calculateIntegrityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIntegrityProcessor) CalculateIntegrityReturns(result1 []byte, result2 error) {
	fake.calculateIntegrityMutex.Lock()
	defer fake.calculateIntegrityMutex.Unlock()
	fake.CalculateIntegrityStub = nil
	fake.calculateIntegrityReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeIntegrityProcessor) CalculateIntegrityReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.calculateIntegrityMutex.Lock()
	defer fake.calculateIntegrityMutex.Unlock()
	fake.CalculateIntegrityStub = nil
	if fake.calculateIntegrityReturnsOnCall == nil {
		fake.calculateIntegrityReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.calculateIntegrityReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeIntegrityProcessor) ValidateIntegrity(arg1 security.IntegralObject) bool {
	fake.validateIntegrityMutex.Lock()
	ret, specificReturn := fake.validateIntegrityReturnsOnCall[len(fake.validateIntegrityArgsForCall)]
	fake.validateIntegrityArgsForCall = append(fake.validateIntegrityArgsForCall, struct {
		arg1 security.IntegralObject
	}{arg1})
	fake.recordInvocation("ValidateIntegrity", []interface{}{arg1})
	fake.validateIntegrityMutex.Unlock()
	if fake.ValidateIntegrityStub != nil {
		return fake.ValidateIntegrityStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateIntegrityReturns
	return fakeReturns.result1
}

func (fake *FakeIntegrityProcessor) ValidateIntegrityCallCount() int {
	fake.validateIntegrityMutex.RLock()
	defer fake.validateIntegrityMutex.RUnlock()
	return len(fake.validateIntegrityArgsForCall)
}

func (fake *FakeIntegrityProcessor) ValidateIntegrityCalls(stub func(security.IntegralObject) bool) {
	fake.validateIntegrityMutex.Lock()
	defer fake.validateIntegrityMutex.Unlock()
	fake.ValidateIntegrityStub = stub
}

func (fake *FakeIntegrityProcessor) ValidateIntegrityArgsForCall(i int) security.IntegralObject {
	fake.validateIntegrityMutex.RLock()
	defer fake.validateIntegrityMutex.RUnlock()
	argsForCall := fake.validateIntegrityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIntegrityProcessor) ValidateIntegrityReturns(result1 bool) {
	fake.validateIntegrityMutex.Lock()
	defer fake.validateIntegrityMutex.Unlock()
	fake.ValidateIntegrityStub = nil
	fake.validateIntegrityReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeIntegrityProcessor) ValidateIntegrityReturnsOnCall(i int, result1 bool) {
	fake.validateIntegrityMutex.Lock()
	defer fake.validateIntegrityMutex.Unlock()
	fake.ValidateIntegrityStub = nil
	if fake.validateIntegrityReturnsOnCall == nil {
		fake.validateIntegrityReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.validateIntegrityReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeIntegrityProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.calculateIntegrityMutex.RLock()
	defer fake.calculateIntegrityMutex.RUnlock()
	fake.validateIntegrityMutex.RLock()
	defer fake.validateIntegrityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIntegrityProcessor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ security.IntegrityProcessor = new(FakeIntegrityProcessor)
