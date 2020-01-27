// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_mc_manager is a generated GoMock package.
package mock_mc_manager

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mc_manager "github.com/solo-io/mesh-projects/services/common/multicluster/manager"
	rest "k8s.io/client-go/rest"
	manager "sigs.k8s.io/controller-runtime/pkg/manager"
)

// MockAsyncManager is a mock of AsyncManager interface
type MockAsyncManager struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerMockRecorder
}

// MockAsyncManagerMockRecorder is the mock recorder for MockAsyncManager
type MockAsyncManagerMockRecorder struct {
	mock *MockAsyncManager
}

// NewMockAsyncManager creates a new mock instance
func NewMockAsyncManager(ctrl *gomock.Controller) *MockAsyncManager {
	mock := &MockAsyncManager{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAsyncManager) EXPECT() *MockAsyncManagerMockRecorder {
	return m.recorder
}

// Manager mocks base method
func (m *MockAsyncManager) Manager() manager.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manager")
	ret0, _ := ret[0].(manager.Manager)
	return ret0
}

// Manager indicates an expected call of Manager
func (mr *MockAsyncManagerMockRecorder) Manager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manager", reflect.TypeOf((*MockAsyncManager)(nil).Manager))
}

// Context mocks base method
func (m *MockAsyncManager) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockAsyncManagerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAsyncManager)(nil).Context))
}

// Error mocks base method
func (m *MockAsyncManager) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockAsyncManagerMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockAsyncManager)(nil).Error))
}

// GotError mocks base method
func (m *MockAsyncManager) GotError() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GotError")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// GotError indicates an expected call of GotError
func (mr *MockAsyncManagerMockRecorder) GotError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GotError", reflect.TypeOf((*MockAsyncManager)(nil).GotError))
}

// Start mocks base method
func (m *MockAsyncManager) Start(opts ...mc_manager.AsyncManagerStartOptionsFunc) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockAsyncManagerMockRecorder) Start(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAsyncManager)(nil).Start), opts...)
}

// Stop mocks base method
func (m *MockAsyncManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockAsyncManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAsyncManager)(nil).Stop))
}

// MockAsyncManagerFactory is a mock of AsyncManagerFactory interface
type MockAsyncManagerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerFactoryMockRecorder
}

// MockAsyncManagerFactoryMockRecorder is the mock recorder for MockAsyncManagerFactory
type MockAsyncManagerFactoryMockRecorder struct {
	mock *MockAsyncManagerFactory
}

// NewMockAsyncManagerFactory creates a new mock instance
func NewMockAsyncManagerFactory(ctrl *gomock.Controller) *MockAsyncManagerFactory {
	mock := &MockAsyncManagerFactory{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAsyncManagerFactory) EXPECT() *MockAsyncManagerFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockAsyncManagerFactory) New(parentCtx context.Context, cfg *rest.Config, opts mc_manager.AsyncManagerOptions) (mc_manager.AsyncManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", parentCtx, cfg, opts)
	ret0, _ := ret[0].(mc_manager.AsyncManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockAsyncManagerFactoryMockRecorder) New(parentCtx, cfg, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockAsyncManagerFactory)(nil).New), parentCtx, cfg, opts)
}
