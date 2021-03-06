// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_compute_target is a generated GoMock package.
package mock_compute_target

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockComputeTargetCredentialsHandler is a mock of ComputeTargetCredentialsHandler interface.
type MockComputeTargetCredentialsHandler struct {
	ctrl     *gomock.Controller
	recorder *MockComputeTargetCredentialsHandlerMockRecorder
}

// MockComputeTargetCredentialsHandlerMockRecorder is the mock recorder for MockComputeTargetCredentialsHandler.
type MockComputeTargetCredentialsHandlerMockRecorder struct {
	mock *MockComputeTargetCredentialsHandler
}

// NewMockComputeTargetCredentialsHandler creates a new mock instance.
func NewMockComputeTargetCredentialsHandler(ctrl *gomock.Controller) *MockComputeTargetCredentialsHandler {
	mock := &MockComputeTargetCredentialsHandler{ctrl: ctrl}
	mock.recorder = &MockComputeTargetCredentialsHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeTargetCredentialsHandler) EXPECT() *MockComputeTargetCredentialsHandlerMockRecorder {
	return m.recorder
}

// ComputeTargetAdded mocks base method.
func (m *MockComputeTargetCredentialsHandler) ComputeTargetAdded(ctx context.Context, secret *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeTargetAdded", ctx, secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComputeTargetAdded indicates an expected call of ComputeTargetAdded.
func (mr *MockComputeTargetCredentialsHandlerMockRecorder) ComputeTargetAdded(ctx, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeTargetAdded", reflect.TypeOf((*MockComputeTargetCredentialsHandler)(nil).ComputeTargetAdded), ctx, secret)
}

// ComputeTargetRemoved mocks base method.
func (m *MockComputeTargetCredentialsHandler) ComputeTargetRemoved(ctx context.Context, secret *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeTargetRemoved", ctx, secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComputeTargetRemoved indicates an expected call of ComputeTargetRemoved.
func (mr *MockComputeTargetCredentialsHandlerMockRecorder) ComputeTargetRemoved(ctx, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeTargetRemoved", reflect.TypeOf((*MockComputeTargetCredentialsHandler)(nil).ComputeTargetRemoved), ctx, secret)
}
