// Code generated by MockGen. DO NOT EDIT.
// Source: unit6.go
//
// Generated by this command:
//
//	mockgen -source=unit6.go -destination=./mocks/unit/unit6.gomock.go
//

// Package mock_unit is a generated GoMock package.
package mock_unit

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFooer is a mock of Fooer interface.
type MockFooer struct {
	ctrl     *gomock.Controller
	recorder *MockFooerMockRecorder
}

// MockFooerMockRecorder is the mock recorder for MockFooer.
type MockFooerMockRecorder struct {
	mock *MockFooer
}

// NewMockFooer creates a new mock instance.
func NewMockFooer(ctrl *gomock.Controller) *MockFooer {
	mock := &MockFooer{ctrl: ctrl}
	mock.recorder = &MockFooerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFooer) EXPECT() *MockFooerMockRecorder {
	return m.recorder
}

// Foo mocks base method.
func (m *MockFooer) Foo() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo")
}

// Foo indicates an expected call of Foo.
func (mr *MockFooerMockRecorder) Foo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockFooer)(nil).Foo))
}
