// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/api (interfaces: WithdrawalService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/data-node/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWithdrawalService is a mock of WithdrawalService interface
type MockWithdrawalService struct {
	ctrl     *gomock.Controller
	recorder *MockWithdrawalServiceMockRecorder
}

// MockWithdrawalServiceMockRecorder is the mock recorder for MockWithdrawalService
type MockWithdrawalServiceMockRecorder struct {
	mock *MockWithdrawalService
}

// NewMockWithdrawalService creates a new mock instance
func NewMockWithdrawalService(ctrl *gomock.Controller) *MockWithdrawalService {
	mock := &MockWithdrawalService{ctrl: ctrl}
	mock.recorder = &MockWithdrawalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWithdrawalService) EXPECT() *MockWithdrawalServiceMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockWithdrawalService) GetByID(arg0 string) (proto.Withdrawal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(proto.Withdrawal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockWithdrawalServiceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockWithdrawalService)(nil).GetByID), arg0)
}

// GetByParty mocks base method
func (m *MockWithdrawalService) GetByParty(arg0 string, arg1 bool) []proto.Withdrawal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1)
	ret0, _ := ret[0].([]proto.Withdrawal)
	return ret0
}

// GetByParty indicates an expected call of GetByParty
func (mr *MockWithdrawalServiceMockRecorder) GetByParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockWithdrawalService)(nil).GetByParty), arg0, arg1)
}
