// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: KeysService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKeysService is a mock of KeysService interface
type MockKeysService struct {
	ctrl     *gomock.Controller
	recorder *MockKeysServiceMockRecorder
}

// MockKeysServiceMockRecorder is the mock recorder for MockKeysService
type MockKeysServiceMockRecorder struct {
	mock *MockKeysService
}

// NewMockKeysService creates a new mock instance
func NewMockKeysService(ctrl *gomock.Controller) *MockKeysService {
	mock := &MockKeysService{ctrl: ctrl}
	mock.recorder = &MockKeysServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeysService) EXPECT() *MockKeysServiceMockRecorder {
	return m.recorder
}

// KeysAdd mocks base method
func (m *MockKeysService) KeysAdd(arg0, arg1 string) (*go_scalingo.Key, error) {
	ret := m.ctrl.Call(m, "KeysAdd", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeysAdd indicates an expected call of KeysAdd
func (mr *MockKeysServiceMockRecorder) KeysAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysAdd", reflect.TypeOf((*MockKeysService)(nil).KeysAdd), arg0, arg1)
}

// KeysDelete mocks base method
func (m *MockKeysService) KeysDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "KeysDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// KeysDelete indicates an expected call of KeysDelete
func (mr *MockKeysServiceMockRecorder) KeysDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysDelete", reflect.TypeOf((*MockKeysService)(nil).KeysDelete), arg0)
}

// KeysList mocks base method
func (m *MockKeysService) KeysList() ([]go_scalingo.Key, error) {
	ret := m.ctrl.Call(m, "KeysList")
	ret0, _ := ret[0].([]go_scalingo.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeysList indicates an expected call of KeysList
func (mr *MockKeysServiceMockRecorder) KeysList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysList", reflect.TypeOf((*MockKeysService)(nil).KeysList))
}
