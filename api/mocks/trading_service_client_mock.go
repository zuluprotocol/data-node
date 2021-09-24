// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/api (interfaces: TradingServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	api "code.vegaprotocol.io/protos/vega/api"
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTradingServiceClient is a mock of TradingServiceClient interface
type MockTradingServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTradingServiceClientMockRecorder
}

// MockTradingServiceClientMockRecorder is the mock recorder for MockTradingServiceClient
type MockTradingServiceClientMockRecorder struct {
	mock *MockTradingServiceClient
}

// NewMockTradingServiceClient creates a new mock instance
func NewMockTradingServiceClient(ctrl *gomock.Controller) *MockTradingServiceClient {
	mock := &MockTradingServiceClient{ctrl: ctrl}
	mock.recorder = &MockTradingServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradingServiceClient) EXPECT() *MockTradingServiceClientMockRecorder {
	return m.recorder
}

// GetVegaTime mocks base method
func (m *MockTradingServiceClient) GetVegaTime(arg0 context.Context, arg1 *api.GetVegaTimeRequest, arg2 ...grpc.CallOption) (*api.GetVegaTimeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVegaTime", varargs...)
	ret0, _ := ret[0].(*api.GetVegaTimeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVegaTime indicates an expected call of GetVegaTime
func (mr *MockTradingServiceClientMockRecorder) GetVegaTime(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVegaTime", reflect.TypeOf((*MockTradingServiceClient)(nil).GetVegaTime), varargs...)
}

// LastBlockHeight mocks base method
func (m *MockTradingServiceClient) LastBlockHeight(arg0 context.Context, arg1 *api.LastBlockHeightRequest, arg2 ...grpc.CallOption) (*api.LastBlockHeightResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LastBlockHeight", varargs...)
	ret0, _ := ret[0].(*api.LastBlockHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastBlockHeight indicates an expected call of LastBlockHeight
func (mr *MockTradingServiceClientMockRecorder) LastBlockHeight(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockHeight", reflect.TypeOf((*MockTradingServiceClient)(nil).LastBlockHeight), varargs...)
}

// ObserveEventBus mocks base method
func (m *MockTradingServiceClient) ObserveEventBus(arg0 context.Context, arg1 ...grpc.CallOption) (api.TradingService_ObserveEventBusClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObserveEventBus", varargs...)
	ret0, _ := ret[0].(api.TradingService_ObserveEventBusClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObserveEventBus indicates an expected call of ObserveEventBus
func (mr *MockTradingServiceClientMockRecorder) ObserveEventBus(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveEventBus", reflect.TypeOf((*MockTradingServiceClient)(nil).ObserveEventBus), varargs...)
}

// PropagateChainEvent mocks base method
func (m *MockTradingServiceClient) PropagateChainEvent(arg0 context.Context, arg1 *api.PropagateChainEventRequest, arg2 ...grpc.CallOption) (*api.PropagateChainEventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PropagateChainEvent", varargs...)
	ret0, _ := ret[0].(*api.PropagateChainEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PropagateChainEvent indicates an expected call of PropagateChainEvent
func (mr *MockTradingServiceClientMockRecorder) PropagateChainEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PropagateChainEvent", reflect.TypeOf((*MockTradingServiceClient)(nil).PropagateChainEvent), varargs...)
}

// Statistics mocks base method
func (m *MockTradingServiceClient) Statistics(arg0 context.Context, arg1 *api.StatisticsRequest, arg2 ...grpc.CallOption) (*api.StatisticsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Statistics", varargs...)
	ret0, _ := ret[0].(*api.StatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Statistics indicates an expected call of Statistics
func (mr *MockTradingServiceClientMockRecorder) Statistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statistics", reflect.TypeOf((*MockTradingServiceClient)(nil).Statistics), varargs...)
}

// SubmitTransaction mocks base method
func (m *MockTradingServiceClient) SubmitTransaction(arg0 context.Context, arg1 *api.SubmitTransactionRequest, arg2 ...grpc.CallOption) (*api.SubmitTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitTransaction", varargs...)
	ret0, _ := ret[0].(*api.SubmitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction
func (mr *MockTradingServiceClientMockRecorder) SubmitTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockTradingServiceClient)(nil).SubmitTransaction), varargs...)
}
