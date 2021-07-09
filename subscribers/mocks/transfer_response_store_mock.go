// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/subscribers (interfaces: TransferResponseStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/data-node/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransferResponseStore is a mock of TransferResponseStore interface
type MockTransferResponseStore struct {
	ctrl     *gomock.Controller
	recorder *MockTransferResponseStoreMockRecorder
}

// MockTransferResponseStoreMockRecorder is the mock recorder for MockTransferResponseStore
type MockTransferResponseStoreMockRecorder struct {
	mock *MockTransferResponseStore
}

// NewMockTransferResponseStore creates a new mock instance
func NewMockTransferResponseStore(ctrl *gomock.Controller) *MockTransferResponseStore {
	mock := &MockTransferResponseStore{ctrl: ctrl}
	mock.recorder = &MockTransferResponseStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransferResponseStore) EXPECT() *MockTransferResponseStoreMockRecorder {
	return m.recorder
}

// SaveBatch mocks base method
func (m *MockTransferResponseStore) SaveBatch(arg0 []*proto.TransferResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBatch indicates an expected call of SaveBatch
func (mr *MockTransferResponseStoreMockRecorder) SaveBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBatch", reflect.TypeOf((*MockTransferResponseStore)(nil).SaveBatch), arg0)
}
