// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	indexer "github.com/sp0x/torrentd/indexer"
	search "github.com/sp0x/torrentd/indexer/search"
	torznab "github.com/sp0x/torrentd/torznab"
	io "io"
	http "net/http"
	reflect "reflect"
)

// MockInfo is a mock of Info interface
type MockInfo struct {
	ctrl     *gomock.Controller
	recorder *MockInfoMockRecorder
}

// MockInfoMockRecorder is the mock recorder for MockInfo
type MockInfoMockRecorder struct {
	mock *MockInfo
}

// NewMockInfo creates a new mock instance
func NewMockInfo(ctrl *gomock.Controller) *MockInfo {
	mock := &MockInfo{ctrl: ctrl}
	mock.recorder = &MockInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfo) EXPECT() *MockInfoMockRecorder {
	return m.recorder
}

// GetId mocks base method
func (m *MockInfo) GetId() string {

	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId
func (mr *MockInfoMockRecorder) GetId() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockInfo)(nil).GetId))
}

// GetTitle mocks base method
func (m *MockInfo) GetTitle() string {

	ret := m.ctrl.Call(m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTitle indicates an expected call of GetTitle
func (mr *MockInfoMockRecorder) GetTitle() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockInfo)(nil).GetTitle))
}

// GetLanguage mocks base method
func (m *MockInfo) GetLanguage() string {

	ret := m.ctrl.Call(m, "GetLanguage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLanguage indicates an expected call of GetLanguage
func (mr *MockInfoMockRecorder) GetLanguage() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLanguage", reflect.TypeOf((*MockInfo)(nil).GetLanguage))
}

// GetLink mocks base method
func (m *MockInfo) GetLink() string {

	ret := m.ctrl.Call(m, "GetLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLink indicates an expected call of GetLink
func (mr *MockInfoMockRecorder) GetLink() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*MockInfo)(nil).GetLink))
}

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockIndexer) Info() indexer.Info {

	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(indexer.Info)
	return ret0
}

// Info indicates an expected call of Info
func (mr *MockIndexerMockRecorder) Info() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockIndexer)(nil).Info))
}

// Search mocks base method
func (m *MockIndexer) Search(query *torznab.Query, srch search.Instance) (search.Instance, error) {

	ret := m.ctrl.Call(m, "Search", query, srch)
	ret0, _ := ret[0].(search.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIndexerMockRecorder) Search(query, srch interface{}) *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), query, srch)
}

// Download mocks base method
func (m *MockIndexer) Download(urlStr string) (io.ReadCloser, error) {

	ret := m.ctrl.Call(m, "Download", urlStr)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockIndexerMockRecorder) Download(urlStr interface{}) *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockIndexer)(nil).Download), urlStr)
}

// Capabilities mocks base method
func (m *MockIndexer) Capabilities() torznab.Capabilities {

	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(torznab.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities
func (mr *MockIndexerMockRecorder) Capabilities() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockIndexer)(nil).Capabilities))
}

// GetEncoding mocks base method
func (m *MockIndexer) GetEncoding() string {

	ret := m.ctrl.Call(m, "GetEncoding")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEncoding indicates an expected call of GetEncoding
func (mr *MockIndexerMockRecorder) GetEncoding() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncoding", reflect.TypeOf((*MockIndexer)(nil).GetEncoding))
}

// ProcessRequest mocks base method
func (m *MockIndexer) ProcessRequest(req *http.Request) (*http.Response, error) {

	ret := m.ctrl.Call(m, "ProcessRequest", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessRequest indicates an expected call of ProcessRequest
func (mr *MockIndexerMockRecorder) ProcessRequest(req interface{}) *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessRequest", reflect.TypeOf((*MockIndexer)(nil).ProcessRequest), req)
}

// Open mocks base method
func (m *MockIndexer) Open(s *search.ExternalResultItem) (io.ReadCloser, error) {

	ret := m.ctrl.Call(m, "Open", s)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockIndexerMockRecorder) Open(s interface{}) *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockIndexer)(nil).Open), s)
}

// Check mocks base method
func (m *MockIndexer) Check() error {

	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockIndexerMockRecorder) Check() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockIndexer)(nil).Check))
}

// MaxSearchPages mocks base method
func (m *MockIndexer) MaxSearchPages() uint {

	ret := m.ctrl.Call(m, "MaxSearchPages")
	ret0, _ := ret[0].(uint)
	return ret0
}

// MaxSearchPages indicates an expected call of MaxSearchPages
func (mr *MockIndexerMockRecorder) MaxSearchPages() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSearchPages", reflect.TypeOf((*MockIndexer)(nil).MaxSearchPages))
}

// SearchIsSinglePaged mocks base method
func (m *MockIndexer) SearchIsSinglePaged() bool {

	ret := m.ctrl.Call(m, "SearchIsSinglePaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SearchIsSinglePaged indicates an expected call of SearchIsSinglePaged
func (mr *MockIndexerMockRecorder) SearchIsSinglePaged() *gomock.Call {

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIsSinglePaged", reflect.TypeOf((*MockIndexer)(nil).SearchIsSinglePaged))
}
