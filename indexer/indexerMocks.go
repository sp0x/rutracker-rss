// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package indexer is a generated GoMock package.
package indexer

import (
	gomock "github.com/golang/mock/gomock"
	search "github.com/sp0x/torrentd/indexer/search"
	storage "github.com/sp0x/torrentd/storage"
	torznab "github.com/sp0x/torrentd/torznab"
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

// GetID mocks base method
func (m *MockInfo) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockInfoMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockInfo)(nil).GetID))
}

// GetTitle mocks base method
func (m *MockInfo) GetTitle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTitle indicates an expected call of GetTitle
func (mr *MockInfoMockRecorder) GetTitle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockInfo)(nil).GetTitle))
}

// GetLanguage mocks base method
func (m *MockInfo) GetLanguage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLanguage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLanguage indicates an expected call of GetLanguage
func (mr *MockInfoMockRecorder) GetLanguage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLanguage", reflect.TypeOf((*MockInfo)(nil).GetLanguage))
}

// GetLink mocks base method
func (m *MockInfo) GetLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLink indicates an expected call of GetLink
func (mr *MockInfoMockRecorder) GetLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
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
func (m *MockIndexer) Info() Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(Info)
	return ret0
}

// Info indicates an expected call of Info
func (mr *MockIndexerMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockIndexer)(nil).Info))
}

// GetDefinition mocks base method
func (m *MockIndexer) GetDefinition() *Definition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefinition")
	ret0, _ := ret[0].(*Definition)
	return ret0
}

// GetDefinition indicates an expected call of GetDefinition
func (mr *MockIndexerMockRecorder) GetDefinition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefinition", reflect.TypeOf((*MockIndexer)(nil).GetDefinition))
}

// Search mocks base method
func (m *MockIndexer) Search(query *search.Query, srch search.Instance) (search.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query, srch)
	ret0, _ := ret[0].(search.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIndexerMockRecorder) Search(query, srch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), query, srch)
}

// Download mocks base method
func (m *MockIndexer) Download(urlStr string) (*ResponseProxy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", urlStr)
	ret0, _ := ret[0].(*ResponseProxy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockIndexerMockRecorder) Download(urlStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockIndexer)(nil).Download), urlStr)
}

// Capabilities mocks base method
func (m *MockIndexer) Capabilities() torznab.Capabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(torznab.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities
func (mr *MockIndexerMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockIndexer)(nil).Capabilities))
}

// GetEncoding mocks base method
func (m *MockIndexer) GetEncoding() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEncoding")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEncoding indicates an expected call of GetEncoding
func (mr *MockIndexerMockRecorder) GetEncoding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncoding", reflect.TypeOf((*MockIndexer)(nil).GetEncoding))
}

// Open mocks base method
func (m *MockIndexer) Open(s search.ResultItemBase) (*ResponseProxy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", s)
	ret0, _ := ret[0].(*ResponseProxy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockIndexerMockRecorder) Open(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockIndexer)(nil).Open), s)
}

// HealthCheck mocks base method
func (m *MockIndexer) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockIndexerMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockIndexer)(nil).HealthCheck))
}

// MaxSearchPages mocks base method
func (m *MockIndexer) MaxSearchPages() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSearchPages")
	ret0, _ := ret[0].(uint)
	return ret0
}

// MaxSearchPages indicates an expected call of MaxSearchPages
func (mr *MockIndexerMockRecorder) MaxSearchPages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSearchPages", reflect.TypeOf((*MockIndexer)(nil).MaxSearchPages))
}

// SearchIsSinglePaged mocks base method
func (m *MockIndexer) SearchIsSinglePaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIsSinglePaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SearchIsSinglePaged indicates an expected call of SearchIsSinglePaged
func (mr *MockIndexerMockRecorder) SearchIsSinglePaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIsSinglePaged", reflect.TypeOf((*MockIndexer)(nil).SearchIsSinglePaged))
}

// Errors mocks base method
func (m *MockIndexer) Errors() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Errors indicates an expected call of Errors
func (mr *MockIndexerMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*MockIndexer)(nil).Errors))
}

// GetStorage mocks base method
func (m *MockIndexer) GetStorage() storage.ItemStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorage")
	ret0, _ := ret[0].(storage.ItemStorage)
	return ret0
}

// GetStorage indicates an expected call of GetStorage
func (mr *MockIndexerMockRecorder) GetStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorage", reflect.TypeOf((*MockIndexer)(nil).GetStorage))
}

// Site mocks base method
func (m *MockIndexer) Site() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Site")
	ret0, _ := ret[0].(string)
	return ret0
}

// Site indicates an expected call of Site
func (mr *MockIndexerMockRecorder) Site() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Site", reflect.TypeOf((*MockIndexer)(nil).Site))
}