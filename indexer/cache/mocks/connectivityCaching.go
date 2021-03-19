// Code generated by MockGen. DO NOT EDIT.
// Source: connectivityCaching.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConnectivityTester is a mock of ConnectivityTester interface
type MockConnectivityTester struct {
	ctrl     *gomock.Controller
	recorder *MockConnectivityTesterMockRecorder
}

// MockConnectivityTesterMockRecorder is the mock recorder for MockConnectivityTester
type MockConnectivityTesterMockRecorder struct {
	mock *MockConnectivityTester
}

// NewMockConnectivityTester creates a new mock instance
func NewMockConnectivityTester(ctrl *gomock.Controller) *MockConnectivityTester {
	mock := &MockConnectivityTester{ctrl: ctrl}
	mock.recorder = &MockConnectivityTesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnectivityTester) EXPECT() *MockConnectivityTesterMockRecorder {
	return m.recorder
}

// IsOkAndSet mocks base method
func (m *MockConnectivityTester) IsOkAndSet(testURL string, f func() bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOkAndSet", testURL, f)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOkAndSet indicates an expected call of IsOkAndSet
func (mr *MockConnectivityTesterMockRecorder) IsOkAndSet(testURL, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOkAndSet", reflect.TypeOf((*MockConnectivityTester)(nil).IsOkAndSet), testURL, f)
}

// IsOk mocks base method
func (m *MockConnectivityTester) IsOk(testURL string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOk", testURL)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOk indicates an expected call of IsOk
func (mr *MockConnectivityTesterMockRecorder) IsOk(testURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOk", reflect.TypeOf((*MockConnectivityTester)(nil).IsOk), testURL)
}

// Test mocks base method
func (m *MockConnectivityTester) Test(testURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Test", testURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// Test indicates an expected call of Test
func (mr *MockConnectivityTesterMockRecorder) Test(testURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockConnectivityTester)(nil).Test), testURL)
}

// Invalidate mocks base method
func (m *MockConnectivityTester) Invalidate(url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Invalidate", url)
}

// Invalidate indicates an expected call of Invalidate
func (mr *MockConnectivityTesterMockRecorder) Invalidate(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockConnectivityTester)(nil).Invalidate), url)
}
