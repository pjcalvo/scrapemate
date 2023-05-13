// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gosom/scrapemate (interfaces: HTTPFetcher)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scrapemate "github.com/gosom/scrapemate"
)

// MockHTTPFetcher is a mock of HTTPFetcher interface.
type MockHTTPFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPFetcherMockRecorder
}

// MockHTTPFetcherMockRecorder is the mock recorder for MockHTTPFetcher.
type MockHTTPFetcherMockRecorder struct {
	mock *MockHTTPFetcher
}

// NewMockHTTPFetcher creates a new mock instance.
func NewMockHTTPFetcher(ctrl *gomock.Controller) *MockHTTPFetcher {
	mock := &MockHTTPFetcher{ctrl: ctrl}
	mock.recorder = &MockHTTPFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPFetcher) EXPECT() *MockHTTPFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockHTTPFetcher) Fetch(arg0 context.Context, arg1 scrapemate.IJob) scrapemate.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1)
	ret0, _ := ret[0].(scrapemate.Response)
	return ret0
}

// Fetch indicates an expected call of Fetch.
func (mr *MockHTTPFetcherMockRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockHTTPFetcher)(nil).Fetch), arg0, arg1)
}
