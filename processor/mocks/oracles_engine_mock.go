// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: OraclesEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	oracles "code.vegaprotocol.io/data-node/oracles"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOraclesEngine is a mock of OraclesEngine interface
type MockOraclesEngine struct {
	ctrl     *gomock.Controller
	recorder *MockOraclesEngineMockRecorder
}

// MockOraclesEngineMockRecorder is the mock recorder for MockOraclesEngine
type MockOraclesEngineMockRecorder struct {
	mock *MockOraclesEngine
}

// NewMockOraclesEngine creates a new mock instance
func NewMockOraclesEngine(ctrl *gomock.Controller) *MockOraclesEngine {
	mock := &MockOraclesEngine{ctrl: ctrl}
	mock.recorder = &MockOraclesEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOraclesEngine) EXPECT() *MockOraclesEngineMockRecorder {
	return m.recorder
}

// BroadcastData mocks base method
func (m *MockOraclesEngine) BroadcastData(arg0 context.Context, arg1 oracles.OracleData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastData indicates an expected call of BroadcastData
func (mr *MockOraclesEngineMockRecorder) BroadcastData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastData", reflect.TypeOf((*MockOraclesEngine)(nil).BroadcastData), arg0, arg1)
}
