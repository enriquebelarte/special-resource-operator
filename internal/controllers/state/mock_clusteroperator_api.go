// Code generated by MockGen. DO NOT EDIT.
// Source: clusteroperator.go

// Package state is a generated GoMock package.
package state

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/config/v1"
)

// MockClusterOperatorManager is a mock of ClusterOperatorManager interface.
type MockClusterOperatorManager struct {
	ctrl     *gomock.Controller
	recorder *MockClusterOperatorManagerMockRecorder
}

// MockClusterOperatorManagerMockRecorder is the mock recorder for MockClusterOperatorManager.
type MockClusterOperatorManagerMockRecorder struct {
	mock *MockClusterOperatorManager
}

// NewMockClusterOperatorManager creates a new mock instance.
func NewMockClusterOperatorManager(ctrl *gomock.Controller) *MockClusterOperatorManager {
	mock := &MockClusterOperatorManager{ctrl: ctrl}
	mock.recorder = &MockClusterOperatorManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterOperatorManager) EXPECT() *MockClusterOperatorManagerMockRecorder {
	return m.recorder
}

// GetOrCreate mocks base method.
func (m *MockClusterOperatorManager) GetOrCreate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOrCreate indicates an expected call of GetOrCreate.
func (mr *MockClusterOperatorManagerMockRecorder) GetOrCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreate", reflect.TypeOf((*MockClusterOperatorManager)(nil).GetOrCreate), arg0)
}

// Refresh mocks base method.
func (m *MockClusterOperatorManager) Refresh(arg0 context.Context, arg1 []v1.ClusterOperatorStatusCondition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh.
func (mr *MockClusterOperatorManagerMockRecorder) Refresh(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockClusterOperatorManager)(nil).Refresh), arg0, arg1)
}