// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mock_mc_manager is a generated GoMock package.
package mock_mc_manager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mc_manager "github.com/solo-io/mesh-projects/services/common/multicluster/manager"
	rest "k8s.io/client-go/rest"
)

// MockKubeConfigHandler is a mock of KubeConfigHandler interface
type MockKubeConfigHandler struct {
	ctrl     *gomock.Controller
	recorder *MockKubeConfigHandlerMockRecorder
}

// MockKubeConfigHandlerMockRecorder is the mock recorder for MockKubeConfigHandler
type MockKubeConfigHandlerMockRecorder struct {
	mock *MockKubeConfigHandler
}

// NewMockKubeConfigHandler creates a new mock instance
func NewMockKubeConfigHandler(ctrl *gomock.Controller) *MockKubeConfigHandler {
	mock := &MockKubeConfigHandler{ctrl: ctrl}
	mock.recorder = &MockKubeConfigHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeConfigHandler) EXPECT() *MockKubeConfigHandlerMockRecorder {
	return m.recorder
}

// ClusterAdded mocks base method
func (m *MockKubeConfigHandler) ClusterAdded(cfg *rest.Config, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterAdded", cfg, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterAdded indicates an expected call of ClusterAdded
func (mr *MockKubeConfigHandlerMockRecorder) ClusterAdded(cfg, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAdded", reflect.TypeOf((*MockKubeConfigHandler)(nil).ClusterAdded), cfg, name)
}

// ClusterRemoved mocks base method
func (m *MockKubeConfigHandler) ClusterRemoved(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRemoved", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterRemoved indicates an expected call of ClusterRemoved
func (mr *MockKubeConfigHandlerMockRecorder) ClusterRemoved(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRemoved", reflect.TypeOf((*MockKubeConfigHandler)(nil).ClusterRemoved), name)
}

// MockAsyncManagerHandler is a mock of AsyncManagerHandler interface
type MockAsyncManagerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerHandlerMockRecorder
}

// MockAsyncManagerHandlerMockRecorder is the mock recorder for MockAsyncManagerHandler
type MockAsyncManagerHandlerMockRecorder struct {
	mock *MockAsyncManagerHandler
}

// NewMockAsyncManagerHandler creates a new mock instance
func NewMockAsyncManagerHandler(ctrl *gomock.Controller) *MockAsyncManagerHandler {
	mock := &MockAsyncManagerHandler{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAsyncManagerHandler) EXPECT() *MockAsyncManagerHandlerMockRecorder {
	return m.recorder
}

// ClusterAdded mocks base method
func (m *MockAsyncManagerHandler) ClusterAdded(mgr mc_manager.AsyncManager, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterAdded", mgr, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterAdded indicates an expected call of ClusterAdded
func (mr *MockAsyncManagerHandlerMockRecorder) ClusterAdded(mgr, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAdded", reflect.TypeOf((*MockAsyncManagerHandler)(nil).ClusterAdded), mgr, name)
}

// ClusterRemoved mocks base method
func (m *MockAsyncManagerHandler) ClusterRemoved(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRemoved", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterRemoved indicates an expected call of ClusterRemoved
func (mr *MockAsyncManagerHandlerMockRecorder) ClusterRemoved(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRemoved", reflect.TypeOf((*MockAsyncManagerHandler)(nil).ClusterRemoved), name)
}

// MockAsyncManagerInformer is a mock of AsyncManagerInformer interface
type MockAsyncManagerInformer struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerInformerMockRecorder
}

// MockAsyncManagerInformerMockRecorder is the mock recorder for MockAsyncManagerInformer
type MockAsyncManagerInformerMockRecorder struct {
	mock *MockAsyncManagerInformer
}

// NewMockAsyncManagerInformer creates a new mock instance
func NewMockAsyncManagerInformer(ctrl *gomock.Controller) *MockAsyncManagerInformer {
	mock := &MockAsyncManagerInformer{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAsyncManagerInformer) EXPECT() *MockAsyncManagerInformerMockRecorder {
	return m.recorder
}

// AddHandler mocks base method
func (m *MockAsyncManagerInformer) AddHandler(informer mc_manager.AsyncManagerHandler, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHandler", informer, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHandler indicates an expected call of AddHandler
func (mr *MockAsyncManagerInformerMockRecorder) AddHandler(informer, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockAsyncManagerInformer)(nil).AddHandler), informer, name)
}

// RemoveHandler mocks base method
func (m *MockAsyncManagerInformer) RemoveHandler(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveHandler", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveHandler indicates an expected call of RemoveHandler
func (mr *MockAsyncManagerInformerMockRecorder) RemoveHandler(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHandler", reflect.TypeOf((*MockAsyncManagerInformer)(nil).RemoveHandler), name)
}
