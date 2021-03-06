// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: TokenGenerator)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTokenGenerator is a mock of TokenGenerator interface
type MockTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenGeneratorMockRecorder
}

// MockTokenGeneratorMockRecorder is the mock recorder for MockTokenGenerator
type MockTokenGeneratorMockRecorder struct {
	mock *MockTokenGenerator
}

// NewMockTokenGenerator creates a new mock instance
func NewMockTokenGenerator(ctrl *gomock.Controller) *MockTokenGenerator {
	mock := &MockTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenGenerator) EXPECT() *MockTokenGeneratorMockRecorder {
	return m.recorder
}

// GetAccessToken mocks base method
func (m *MockTokenGenerator) GetAccessToken() (string, error) {
	ret := m.ctrl.Call(m, "GetAccessToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessToken indicates an expected call of GetAccessToken
func (mr *MockTokenGeneratorMockRecorder) GetAccessToken() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockTokenGenerator)(nil).GetAccessToken))
}

// SetClient mocks base method
func (m *MockTokenGenerator) SetClient(arg0 *go_scalingo.Client) {
	m.ctrl.Call(m, "SetClient", arg0)
}

// SetClient indicates an expected call of SetClient
func (mr *MockTokenGeneratorMockRecorder) SetClient(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockTokenGenerator)(nil).SetClient), arg0)
}
