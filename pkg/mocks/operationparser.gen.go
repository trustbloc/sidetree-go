// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-go/pkg/api/operation"
)

type OperationParser struct {
	GetCommitmentStub        func([]byte) (string, error)
	getCommitmentMutex       sync.RWMutex
	getCommitmentArgsForCall []struct {
		arg1 []byte
	}
	getCommitmentReturns struct {
		result1 string
		result2 error
	}
	getCommitmentReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRevealValueStub        func([]byte) (string, error)
	getRevealValueMutex       sync.RWMutex
	getRevealValueArgsForCall []struct {
		arg1 []byte
	}
	getRevealValueReturns struct {
		result1 string
		result2 error
	}
	getRevealValueReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ParseStub        func(string, []byte) (*operation.Operation, error)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	parseReturns struct {
		result1 *operation.Operation
		result2 error
	}
	parseReturnsOnCall map[int]struct {
		result1 *operation.Operation
		result2 error
	}
	ParseDIDStub        func(string, string) (string, []byte, error)
	parseDIDMutex       sync.RWMutex
	parseDIDArgsForCall []struct {
		arg1 string
		arg2 string
	}
	parseDIDReturns struct {
		result1 string
		result2 []byte
		result3 error
	}
	parseDIDReturnsOnCall map[int]struct {
		result1 string
		result2 []byte
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OperationParser) GetCommitment(arg1 []byte) (string, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getCommitmentMutex.Lock()
	ret, specificReturn := fake.getCommitmentReturnsOnCall[len(fake.getCommitmentArgsForCall)]
	fake.getCommitmentArgsForCall = append(fake.getCommitmentArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.GetCommitmentStub
	fakeReturns := fake.getCommitmentReturns
	fake.recordInvocation("GetCommitment", []interface{}{arg1Copy})
	fake.getCommitmentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OperationParser) GetCommitmentCallCount() int {
	fake.getCommitmentMutex.RLock()
	defer fake.getCommitmentMutex.RUnlock()
	return len(fake.getCommitmentArgsForCall)
}

func (fake *OperationParser) GetCommitmentCalls(stub func([]byte) (string, error)) {
	fake.getCommitmentMutex.Lock()
	defer fake.getCommitmentMutex.Unlock()
	fake.GetCommitmentStub = stub
}

func (fake *OperationParser) GetCommitmentArgsForCall(i int) []byte {
	fake.getCommitmentMutex.RLock()
	defer fake.getCommitmentMutex.RUnlock()
	argsForCall := fake.getCommitmentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *OperationParser) GetCommitmentReturns(result1 string, result2 error) {
	fake.getCommitmentMutex.Lock()
	defer fake.getCommitmentMutex.Unlock()
	fake.GetCommitmentStub = nil
	fake.getCommitmentReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OperationParser) GetCommitmentReturnsOnCall(i int, result1 string, result2 error) {
	fake.getCommitmentMutex.Lock()
	defer fake.getCommitmentMutex.Unlock()
	fake.GetCommitmentStub = nil
	if fake.getCommitmentReturnsOnCall == nil {
		fake.getCommitmentReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getCommitmentReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OperationParser) GetRevealValue(arg1 []byte) (string, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getRevealValueMutex.Lock()
	ret, specificReturn := fake.getRevealValueReturnsOnCall[len(fake.getRevealValueArgsForCall)]
	fake.getRevealValueArgsForCall = append(fake.getRevealValueArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.GetRevealValueStub
	fakeReturns := fake.getRevealValueReturns
	fake.recordInvocation("GetRevealValue", []interface{}{arg1Copy})
	fake.getRevealValueMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OperationParser) GetRevealValueCallCount() int {
	fake.getRevealValueMutex.RLock()
	defer fake.getRevealValueMutex.RUnlock()
	return len(fake.getRevealValueArgsForCall)
}

func (fake *OperationParser) GetRevealValueCalls(stub func([]byte) (string, error)) {
	fake.getRevealValueMutex.Lock()
	defer fake.getRevealValueMutex.Unlock()
	fake.GetRevealValueStub = stub
}

func (fake *OperationParser) GetRevealValueArgsForCall(i int) []byte {
	fake.getRevealValueMutex.RLock()
	defer fake.getRevealValueMutex.RUnlock()
	argsForCall := fake.getRevealValueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *OperationParser) GetRevealValueReturns(result1 string, result2 error) {
	fake.getRevealValueMutex.Lock()
	defer fake.getRevealValueMutex.Unlock()
	fake.GetRevealValueStub = nil
	fake.getRevealValueReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OperationParser) GetRevealValueReturnsOnCall(i int, result1 string, result2 error) {
	fake.getRevealValueMutex.Lock()
	defer fake.getRevealValueMutex.Unlock()
	fake.GetRevealValueStub = nil
	if fake.getRevealValueReturnsOnCall == nil {
		fake.getRevealValueReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getRevealValueReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OperationParser) Parse(arg1 string, arg2 []byte) (*operation.Operation, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.parseMutex.Lock()
	ret, specificReturn := fake.parseReturnsOnCall[len(fake.parseArgsForCall)]
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.ParseStub
	fakeReturns := fake.parseReturns
	fake.recordInvocation("Parse", []interface{}{arg1, arg2Copy})
	fake.parseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OperationParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *OperationParser) ParseCalls(stub func(string, []byte) (*operation.Operation, error)) {
	fake.parseMutex.Lock()
	defer fake.parseMutex.Unlock()
	fake.ParseStub = stub
}

func (fake *OperationParser) ParseArgsForCall(i int) (string, []byte) {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	argsForCall := fake.parseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *OperationParser) ParseReturns(result1 *operation.Operation, result2 error) {
	fake.parseMutex.Lock()
	defer fake.parseMutex.Unlock()
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 *operation.Operation
		result2 error
	}{result1, result2}
}

func (fake *OperationParser) ParseReturnsOnCall(i int, result1 *operation.Operation, result2 error) {
	fake.parseMutex.Lock()
	defer fake.parseMutex.Unlock()
	fake.ParseStub = nil
	if fake.parseReturnsOnCall == nil {
		fake.parseReturnsOnCall = make(map[int]struct {
			result1 *operation.Operation
			result2 error
		})
	}
	fake.parseReturnsOnCall[i] = struct {
		result1 *operation.Operation
		result2 error
	}{result1, result2}
}

func (fake *OperationParser) ParseDID(arg1 string, arg2 string) (string, []byte, error) {
	fake.parseDIDMutex.Lock()
	ret, specificReturn := fake.parseDIDReturnsOnCall[len(fake.parseDIDArgsForCall)]
	fake.parseDIDArgsForCall = append(fake.parseDIDArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ParseDIDStub
	fakeReturns := fake.parseDIDReturns
	fake.recordInvocation("ParseDID", []interface{}{arg1, arg2})
	fake.parseDIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *OperationParser) ParseDIDCallCount() int {
	fake.parseDIDMutex.RLock()
	defer fake.parseDIDMutex.RUnlock()
	return len(fake.parseDIDArgsForCall)
}

func (fake *OperationParser) ParseDIDCalls(stub func(string, string) (string, []byte, error)) {
	fake.parseDIDMutex.Lock()
	defer fake.parseDIDMutex.Unlock()
	fake.ParseDIDStub = stub
}

func (fake *OperationParser) ParseDIDArgsForCall(i int) (string, string) {
	fake.parseDIDMutex.RLock()
	defer fake.parseDIDMutex.RUnlock()
	argsForCall := fake.parseDIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *OperationParser) ParseDIDReturns(result1 string, result2 []byte, result3 error) {
	fake.parseDIDMutex.Lock()
	defer fake.parseDIDMutex.Unlock()
	fake.ParseDIDStub = nil
	fake.parseDIDReturns = struct {
		result1 string
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *OperationParser) ParseDIDReturnsOnCall(i int, result1 string, result2 []byte, result3 error) {
	fake.parseDIDMutex.Lock()
	defer fake.parseDIDMutex.Unlock()
	fake.ParseDIDStub = nil
	if fake.parseDIDReturnsOnCall == nil {
		fake.parseDIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 []byte
			result3 error
		})
	}
	fake.parseDIDReturnsOnCall[i] = struct {
		result1 string
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *OperationParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCommitmentMutex.RLock()
	defer fake.getCommitmentMutex.RUnlock()
	fake.getRevealValueMutex.RLock()
	defer fake.getRevealValueMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	fake.parseDIDMutex.RLock()
	defer fake.parseDIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OperationParser) recordInvocation(key string, args []interface{}) {
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
