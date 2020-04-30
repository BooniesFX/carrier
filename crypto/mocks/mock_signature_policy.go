// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/IPFS-eX/carrier/crypto (interfaces: SignaturePolicy)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	crypto "github.com/IPFS-eX/carrier/crypto"
)

// MockSignaturePolicy is a mock of SignaturePolicy interface
type MockSignaturePolicy struct {
	ctrl     *gomock.Controller
	recorder *MockSignaturePolicyMockRecorder
}

// MockSignaturePolicyMockRecorder is the mock recorder for MockSignaturePolicy
type MockSignaturePolicyMockRecorder struct {
	mock *MockSignaturePolicy
}

// NewMockSignaturePolicy creates a new mock instance
func NewMockSignaturePolicy(ctrl *gomock.Controller) *MockSignaturePolicy {
	mock := &MockSignaturePolicy{ctrl: ctrl}
	mock.recorder = &MockSignaturePolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSignaturePolicy) EXPECT() *MockSignaturePolicyMockRecorder {
	return m.recorder
}

// GenerateKeys mocks base method
func (m *MockSignaturePolicy) GenerateKeys() ([]byte, []byte, error) {
	ret := m.ctrl.Call(m, "GenerateKeys")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateKeys indicates an expected call of GenerateKeys
func (mr *MockSignaturePolicyMockRecorder) GenerateKeys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKeys", reflect.TypeOf((*MockSignaturePolicy)(nil).GenerateKeys))
}

// PrivateKeySize mocks base method
func (m *MockSignaturePolicy) PrivateKeySize() int {
	ret := m.ctrl.Call(m, "PrivateKeySize")
	ret0, _ := ret[0].(int)
	return ret0
}

// PrivateKeySize indicates an expected call of PrivateKeySize
func (mr *MockSignaturePolicyMockRecorder) PrivateKeySize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateKeySize", reflect.TypeOf((*MockSignaturePolicy)(nil).PrivateKeySize))
}

// PrivateToPublic mocks base method
func (m *MockSignaturePolicy) PrivateToPublic(arg0 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "PrivateToPublic", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateToPublic indicates an expected call of PrivateToPublic
func (mr *MockSignaturePolicyMockRecorder) PrivateToPublic(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateToPublic", reflect.TypeOf((*MockSignaturePolicy)(nil).PrivateToPublic), arg0)
}

// PublicKeySize mocks base method
func (m *MockSignaturePolicy) PublicKeySize() int {
	ret := m.ctrl.Call(m, "PublicKeySize")
	ret0, _ := ret[0].(int)
	return ret0
}

// PublicKeySize indicates an expected call of PublicKeySize
func (mr *MockSignaturePolicyMockRecorder) PublicKeySize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicKeySize", reflect.TypeOf((*MockSignaturePolicy)(nil).PublicKeySize))
}

// RandomKeyPair mocks base method
func (m *MockSignaturePolicy) RandomKeyPair() *crypto.KeyPair {
	ret := m.ctrl.Call(m, "RandomKeyPair")
	ret0, _ := ret[0].(*crypto.KeyPair)
	return ret0
}

// RandomKeyPair indicates an expected call of RandomKeyPair
func (mr *MockSignaturePolicyMockRecorder) RandomKeyPair() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomKeyPair", reflect.TypeOf((*MockSignaturePolicy)(nil).RandomKeyPair))
}

// Sign mocks base method
func (m *MockSignaturePolicy) Sign(arg0, arg1 []byte) []byte {
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Sign indicates an expected call of Sign
func (mr *MockSignaturePolicyMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSignaturePolicy)(nil).Sign), arg0, arg1)
}

// Verify mocks base method
func (m *MockSignaturePolicy) Verify(arg0, arg1, arg2 []byte) bool {
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockSignaturePolicyMockRecorder) Verify(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockSignaturePolicy)(nil).Verify), arg0, arg1, arg2)
}
