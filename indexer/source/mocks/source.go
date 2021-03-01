// Code generated by MockGen. DO NOT EDIT.
// Source: source.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	source "github.com/sp0x/torrentd/indexer/source"
	io "io"
	url "net/url"
	reflect "reflect"
)

// MockFetchResult is a mock of FetchResult interface
type MockFetchResult struct {
	ctrl     *gomock.Controller
	recorder *MockFetchResultMockRecorder
}

// MockFetchResultMockRecorder is the mock recorder for MockFetchResult
type MockFetchResultMockRecorder struct {
	mock *MockFetchResult
}

// NewMockFetchResult creates a new mock instance
func NewMockFetchResult(ctrl *gomock.Controller) *MockFetchResult {
	mock := &MockFetchResult{ctrl: ctrl}
	mock.recorder = &MockFetchResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetchResult) EXPECT() *MockFetchResultMockRecorder {
	return m.recorder
}

// ContentType mocks base method
func (m *MockFetchResult) ContentType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentType indicates an expected call of ContentType
func (mr *MockFetchResultMockRecorder) ContentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentType", reflect.TypeOf((*MockFetchResult)(nil).ContentType))
}

// Encoding mocks base method
func (m *MockFetchResult) Encoding() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encoding")
	ret0, _ := ret[0].(string)
	return ret0
}

// Encoding indicates an expected call of Encoding
func (mr *MockFetchResultMockRecorder) Encoding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encoding", reflect.TypeOf((*MockFetchResult)(nil).Encoding))
}

// MockContentFetcher is a mock of ContentFetcher interface
type MockContentFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockContentFetcherMockRecorder
}

// MockContentFetcherMockRecorder is the mock recorder for MockContentFetcher
type MockContentFetcherMockRecorder struct {
	mock *MockContentFetcher
}

// NewMockContentFetcher creates a new mock instance
func NewMockContentFetcher(ctrl *gomock.Controller) *MockContentFetcher {
	mock := &MockContentFetcher{ctrl: ctrl}
	mock.recorder = &MockContentFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContentFetcher) EXPECT() *MockContentFetcherMockRecorder {
	return m.recorder
}

// Cleanup mocks base method
func (m *MockContentFetcher) Cleanup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cleanup")
}

// Cleanup indicates an expected call of Cleanup
func (mr *MockContentFetcherMockRecorder) Cleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockContentFetcher)(nil).Cleanup))
}

// Fetch mocks base method
func (m *MockContentFetcher) Fetch(target *source.RequestOptions) (source.FetchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", target)
	ret0, _ := ret[0].(source.FetchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockContentFetcherMockRecorder) Fetch(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockContentFetcher)(nil).Fetch), target)
}

// Post mocks base method
func (m *MockContentFetcher) Post(options *source.RequestOptions) (source.FetchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", options)
	ret0, _ := ret[0].(source.FetchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post
func (mr *MockContentFetcherMockRecorder) Post(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockContentFetcher)(nil).Post), options)
}

// URL mocks base method
func (m *MockContentFetcher) URL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockContentFetcherMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockContentFetcher)(nil).URL))
}

// Clone mocks base method
func (m *MockContentFetcher) Clone() source.ContentFetcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(source.ContentFetcher)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockContentFetcherMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockContentFetcher)(nil).Clone))
}

// Open mocks base method
func (m *MockContentFetcher) Open(options *source.RequestOptions) (source.FetchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", options)
	ret0, _ := ret[0].(source.FetchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockContentFetcherMockRecorder) Open(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockContentFetcher)(nil).Open), options)
}

// Download mocks base method
func (m *MockContentFetcher) Download(buffer io.Writer) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", buffer)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockContentFetcherMockRecorder) Download(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockContentFetcher)(nil).Download), buffer)
}

// SetErrorHandler mocks base method
func (m *MockContentFetcher) SetErrorHandler(callback func(*source.RequestOptions)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetErrorHandler", callback)
}

// SetErrorHandler indicates an expected call of SetErrorHandler
func (mr *MockContentFetcherMockRecorder) SetErrorHandler(callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetErrorHandler", reflect.TypeOf((*MockContentFetcher)(nil).SetErrorHandler), callback)
}
