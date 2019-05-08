// Code generated by counterfeiter. DO NOT EDIT.
package resolverfakes

import (
	"sync"

	"github.com/ankeesler/flexec/resolver"
)

type FakeResolver struct {
	ResolveInputStub        func(string) (string, error)
	resolveInputMutex       sync.RWMutex
	resolveInputArgsForCall []struct {
		arg1 string
	}
	resolveInputReturns struct {
		result1 string
		result2 error
	}
	resolveInputReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ResolveOutputStub        func(string) (string, error)
	resolveOutputMutex       sync.RWMutex
	resolveOutputArgsForCall []struct {
		arg1 string
	}
	resolveOutputReturns struct {
		result1 string
		result2 error
	}
	resolveOutputReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ResolveParamStub        func(string) (string, error)
	resolveParamMutex       sync.RWMutex
	resolveParamArgsForCall []struct {
		arg1 string
	}
	resolveParamReturns struct {
		result1 string
		result2 error
	}
	resolveParamReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResolver) ResolveInput(arg1 string) (string, error) {
	fake.resolveInputMutex.Lock()
	ret, specificReturn := fake.resolveInputReturnsOnCall[len(fake.resolveInputArgsForCall)]
	fake.resolveInputArgsForCall = append(fake.resolveInputArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ResolveInput", []interface{}{arg1})
	fake.resolveInputMutex.Unlock()
	if fake.ResolveInputStub != nil {
		return fake.ResolveInputStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveInputReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResolver) ResolveInputCallCount() int {
	fake.resolveInputMutex.RLock()
	defer fake.resolveInputMutex.RUnlock()
	return len(fake.resolveInputArgsForCall)
}

func (fake *FakeResolver) ResolveInputCalls(stub func(string) (string, error)) {
	fake.resolveInputMutex.Lock()
	defer fake.resolveInputMutex.Unlock()
	fake.ResolveInputStub = stub
}

func (fake *FakeResolver) ResolveInputArgsForCall(i int) string {
	fake.resolveInputMutex.RLock()
	defer fake.resolveInputMutex.RUnlock()
	argsForCall := fake.resolveInputArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResolver) ResolveInputReturns(result1 string, result2 error) {
	fake.resolveInputMutex.Lock()
	defer fake.resolveInputMutex.Unlock()
	fake.ResolveInputStub = nil
	fake.resolveInputReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResolver) ResolveInputReturnsOnCall(i int, result1 string, result2 error) {
	fake.resolveInputMutex.Lock()
	defer fake.resolveInputMutex.Unlock()
	fake.ResolveInputStub = nil
	if fake.resolveInputReturnsOnCall == nil {
		fake.resolveInputReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.resolveInputReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResolver) ResolveOutput(arg1 string) (string, error) {
	fake.resolveOutputMutex.Lock()
	ret, specificReturn := fake.resolveOutputReturnsOnCall[len(fake.resolveOutputArgsForCall)]
	fake.resolveOutputArgsForCall = append(fake.resolveOutputArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ResolveOutput", []interface{}{arg1})
	fake.resolveOutputMutex.Unlock()
	if fake.ResolveOutputStub != nil {
		return fake.ResolveOutputStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveOutputReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResolver) ResolveOutputCallCount() int {
	fake.resolveOutputMutex.RLock()
	defer fake.resolveOutputMutex.RUnlock()
	return len(fake.resolveOutputArgsForCall)
}

func (fake *FakeResolver) ResolveOutputCalls(stub func(string) (string, error)) {
	fake.resolveOutputMutex.Lock()
	defer fake.resolveOutputMutex.Unlock()
	fake.ResolveOutputStub = stub
}

func (fake *FakeResolver) ResolveOutputArgsForCall(i int) string {
	fake.resolveOutputMutex.RLock()
	defer fake.resolveOutputMutex.RUnlock()
	argsForCall := fake.resolveOutputArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResolver) ResolveOutputReturns(result1 string, result2 error) {
	fake.resolveOutputMutex.Lock()
	defer fake.resolveOutputMutex.Unlock()
	fake.ResolveOutputStub = nil
	fake.resolveOutputReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResolver) ResolveOutputReturnsOnCall(i int, result1 string, result2 error) {
	fake.resolveOutputMutex.Lock()
	defer fake.resolveOutputMutex.Unlock()
	fake.ResolveOutputStub = nil
	if fake.resolveOutputReturnsOnCall == nil {
		fake.resolveOutputReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.resolveOutputReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResolver) ResolveParam(arg1 string) (string, error) {
	fake.resolveParamMutex.Lock()
	ret, specificReturn := fake.resolveParamReturnsOnCall[len(fake.resolveParamArgsForCall)]
	fake.resolveParamArgsForCall = append(fake.resolveParamArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ResolveParam", []interface{}{arg1})
	fake.resolveParamMutex.Unlock()
	if fake.ResolveParamStub != nil {
		return fake.ResolveParamStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveParamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResolver) ResolveParamCallCount() int {
	fake.resolveParamMutex.RLock()
	defer fake.resolveParamMutex.RUnlock()
	return len(fake.resolveParamArgsForCall)
}

func (fake *FakeResolver) ResolveParamCalls(stub func(string) (string, error)) {
	fake.resolveParamMutex.Lock()
	defer fake.resolveParamMutex.Unlock()
	fake.ResolveParamStub = stub
}

func (fake *FakeResolver) ResolveParamArgsForCall(i int) string {
	fake.resolveParamMutex.RLock()
	defer fake.resolveParamMutex.RUnlock()
	argsForCall := fake.resolveParamArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResolver) ResolveParamReturns(result1 string, result2 error) {
	fake.resolveParamMutex.Lock()
	defer fake.resolveParamMutex.Unlock()
	fake.ResolveParamStub = nil
	fake.resolveParamReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResolver) ResolveParamReturnsOnCall(i int, result1 string, result2 error) {
	fake.resolveParamMutex.Lock()
	defer fake.resolveParamMutex.Unlock()
	fake.ResolveParamStub = nil
	if fake.resolveParamReturnsOnCall == nil {
		fake.resolveParamReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.resolveParamReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResolver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveInputMutex.RLock()
	defer fake.resolveInputMutex.RUnlock()
	fake.resolveOutputMutex.RLock()
	defer fake.resolveOutputMutex.RUnlock()
	fake.resolveParamMutex.RLock()
	defer fake.resolveParamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResolver) recordInvocation(key string, args []interface{}) {
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

var _ resolver.Resolver = new(FakeResolver)