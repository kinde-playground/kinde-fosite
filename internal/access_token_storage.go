// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AccessTokenStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockAccessTokenStorage is a mock of AccessTokenStorage interface.
type MockAccessTokenStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenStorageMockRecorder
}

// MockAccessTokenStorageMockRecorder is the mock recorder for MockAccessTokenStorage.
type MockAccessTokenStorageMockRecorder struct {
	mock *MockAccessTokenStorage
}

// NewMockAccessTokenStorage creates a new mock instance.
func NewMockAccessTokenStorage(ctrl *gomock.Controller) *MockAccessTokenStorage {
	mock := &MockAccessTokenStorage{ctrl: ctrl}
	mock.recorder = &MockAccessTokenStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokenStorage) EXPECT() *MockAccessTokenStorageMockRecorder {
	return m.recorder
}

// CreateAccessTokenSession mocks base method.
func (m *MockAccessTokenStorage) CreateAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessTokenSession indicates an expected call of CreateAccessTokenSession.
func (mr *MockAccessTokenStorageMockRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessTokenSession", reflect.TypeOf((*MockAccessTokenStorage)(nil).CreateAccessTokenSession), arg0, arg1, arg2)
}

// DeleteAccessTokenSession mocks base method.
func (m *MockAccessTokenStorage) DeleteAccessTokenSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessTokenSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessTokenSession indicates an expected call of DeleteAccessTokenSession.
func (mr *MockAccessTokenStorageMockRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessTokenSession", reflect.TypeOf((*MockAccessTokenStorage)(nil).DeleteAccessTokenSession), arg0, arg1)
}

// GetAccessTokenSession mocks base method.
func (m *MockAccessTokenStorage) GetAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessTokenSession indicates an expected call of GetAccessTokenSession.
func (mr *MockAccessTokenStorageMockRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenSession", reflect.TypeOf((*MockAccessTokenStorage)(nil).GetAccessTokenSession), arg0, arg1, arg2)
}
