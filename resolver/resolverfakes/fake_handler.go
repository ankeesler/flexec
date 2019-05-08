// Code generated by counterfeiter. DO NOT EDIT.
package resolverfakes

import (
	"sync"

	"github.com/ankeesler/flexec/resolver"
)

type FakeHandler struct {
	OnResolvedInputStub        func(string, string)
	onResolvedInputMutex       sync.RWMutex
	onResolvedInputArgsForCall []struct {
		arg1 string
		arg2 string
	}
	OnResolvedOutputStub        func(string, string)
	onResolvedOutputMutex       sync.RWMutex
	onResolvedOutputArgsForCall []struct {
		arg1 string
		arg2 string
	}
	OnResolvedParamStub        func(string, string)
	onResolvedParamMutex       sync.RWMutex
	onResolvedParamArgsForCall []struct {
		arg1 string
		arg2 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandler) OnResolvedInput(arg1 string, arg2 string) {
	fake.onResolvedInputMutex.Lock()
	fake.onResolvedInputArgsForCall = append(fake.onResolvedInputArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("OnResolvedInput", []interface{}{arg1, arg2})
	fake.onResolvedInputMutex.Unlock()
	if fake.OnResolvedInputStub != nil {
		fake.OnResolvedInputStub(arg1, arg2)
	}
}

func (fake *FakeHandler) OnResolvedInputCallCount() int {
	fake.onResolvedInputMutex.RLock()
	defer fake.onResolvedInputMutex.RUnlock()
	return len(fake.onResolvedInputArgsForCall)
}

func (fake *FakeHandler) OnResolvedInputCalls(stub func(string, string)) {
	fake.onResolvedInputMutex.Lock()
	defer fake.onResolvedInputMutex.Unlock()
	fake.OnResolvedInputStub = stub
}

func (fake *FakeHandler) OnResolvedInputArgsForCall(i int) (string, string) {
	fake.onResolvedInputMutex.RLock()
	defer fake.onResolvedInputMutex.RUnlock()
	argsForCall := fake.onResolvedInputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) OnResolvedOutput(arg1 string, arg2 string) {
	fake.onResolvedOutputMutex.Lock()
	fake.onResolvedOutputArgsForCall = append(fake.onResolvedOutputArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("OnResolvedOutput", []interface{}{arg1, arg2})
	fake.onResolvedOutputMutex.Unlock()
	if fake.OnResolvedOutputStub != nil {
		fake.OnResolvedOutputStub(arg1, arg2)
	}
}

func (fake *FakeHandler) OnResolvedOutputCallCount() int {
	fake.onResolvedOutputMutex.RLock()
	defer fake.onResolvedOutputMutex.RUnlock()
	return len(fake.onResolvedOutputArgsForCall)
}

func (fake *FakeHandler) OnResolvedOutputCalls(stub func(string, string)) {
	fake.onResolvedOutputMutex.Lock()
	defer fake.onResolvedOutputMutex.Unlock()
	fake.OnResolvedOutputStub = stub
}

func (fake *FakeHandler) OnResolvedOutputArgsForCall(i int) (string, string) {
	fake.onResolvedOutputMutex.RLock()
	defer fake.onResolvedOutputMutex.RUnlock()
	argsForCall := fake.onResolvedOutputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) OnResolvedParam(arg1 string, arg2 string) {
	fake.onResolvedParamMutex.Lock()
	fake.onResolvedParamArgsForCall = append(fake.onResolvedParamArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("OnResolvedParam", []interface{}{arg1, arg2})
	fake.onResolvedParamMutex.Unlock()
	if fake.OnResolvedParamStub != nil {
		fake.OnResolvedParamStub(arg1, arg2)
	}
}

func (fake *FakeHandler) OnResolvedParamCallCount() int {
	fake.onResolvedParamMutex.RLock()
	defer fake.onResolvedParamMutex.RUnlock()
	return len(fake.onResolvedParamArgsForCall)
}

func (fake *FakeHandler) OnResolvedParamCalls(stub func(string, string)) {
	fake.onResolvedParamMutex.Lock()
	defer fake.onResolvedParamMutex.Unlock()
	fake.OnResolvedParamStub = stub
}

func (fake *FakeHandler) OnResolvedParamArgsForCall(i int) (string, string) {
	fake.onResolvedParamMutex.RLock()
	defer fake.onResolvedParamMutex.RUnlock()
	argsForCall := fake.onResolvedParamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onResolvedInputMutex.RLock()
	defer fake.onResolvedInputMutex.RUnlock()
	fake.onResolvedOutputMutex.RLock()
	defer fake.onResolvedOutputMutex.RUnlock()
	fake.onResolvedParamMutex.RLock()
	defer fake.onResolvedParamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandler) recordInvocation(key string, args []interface{}) {
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

var _ resolver.Handler = new(FakeHandler)
