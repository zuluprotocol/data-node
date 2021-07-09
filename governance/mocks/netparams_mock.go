// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/governance (interfaces: NetParams)

// Package mocks is a generated GoMock package.
package mocks

import (
	netparams "code.vegaprotocol.io/data-node/netparams"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockNetParams is a mock of NetParams interface
type MockNetParams struct {
	ctrl     *gomock.Controller
	recorder *MockNetParamsMockRecorder
}

// MockNetParamsMockRecorder is the mock recorder for MockNetParams
type MockNetParamsMockRecorder struct {
	mock *MockNetParams
}

// NewMockNetParams creates a new mock instance
func NewMockNetParams(ctrl *gomock.Controller) *MockNetParams {
	mock := &MockNetParams{ctrl: ctrl}
	mock.recorder = &MockNetParamsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetParams) EXPECT() *MockNetParamsMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockNetParams) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockNetParamsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNetParams)(nil).Get), arg0)
}

// GetDuration mocks base method
func (m *MockNetParams) GetDuration(arg0 string) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDuration", arg0)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDuration indicates an expected call of GetDuration
func (mr *MockNetParamsMockRecorder) GetDuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDuration", reflect.TypeOf((*MockNetParams)(nil).GetDuration), arg0)
}

// GetFloat mocks base method
func (m *MockNetParams) GetFloat(arg0 string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloat", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloat indicates an expected call of GetFloat
func (mr *MockNetParamsMockRecorder) GetFloat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat", reflect.TypeOf((*MockNetParams)(nil).GetFloat), arg0)
}

// GetInt mocks base method
func (m *MockNetParams) GetInt(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt
func (mr *MockNetParamsMockRecorder) GetInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockNetParams)(nil).GetInt), arg0)
}

// GetJSONStruct mocks base method
func (m *MockNetParams) GetJSONStruct(arg0 string, arg1 netparams.Reset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJSONStruct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetJSONStruct indicates an expected call of GetJSONStruct
func (mr *MockNetParamsMockRecorder) GetJSONStruct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJSONStruct", reflect.TypeOf((*MockNetParams)(nil).GetJSONStruct), arg0, arg1)
}

// Update mocks base method
func (m *MockNetParams) Update(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockNetParamsMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNetParams)(nil).Update), arg0, arg1, arg2)
}

// Validate mocks base method
func (m *MockNetParams) Validate(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockNetParamsMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockNetParams)(nil).Validate), arg0, arg1)
}
