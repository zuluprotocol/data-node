// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/datanode/api (interfaces: RewardsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockRewardsService is a mock of RewardsService interface.
type MockRewardsService struct {
	ctrl     *gomock.Controller
	recorder *MockRewardsServiceMockRecorder
}

// MockRewardsServiceMockRecorder is the mock recorder for MockRewardsService.
type MockRewardsServiceMockRecorder struct {
	mock *MockRewardsService
}

// NewMockRewardsService creates a new mock instance.
func NewMockRewardsService(ctrl *gomock.Controller) *MockRewardsService {
	mock := &MockRewardsService{ctrl: ctrl}
	mock.recorder = &MockRewardsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRewardsService) EXPECT() *MockRewardsServiceMockRecorder {
	return m.recorder
}

// GetRewardSummaries mocks base method.
func (m *MockRewardsService) GetRewardSummaries(arg0 context.Context, arg1 string, arg2 *string) []*vega.RewardSummary {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardSummaries", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*vega.RewardSummary)
	return ret0
}

// GetRewardSummaries indicates an expected call of GetRewardSummaries.
func (mr *MockRewardsServiceMockRecorder) GetRewardSummaries(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardSummaries", reflect.TypeOf((*MockRewardsService)(nil).GetRewardSummaries), arg0, arg1, arg2)
}

// GetRewards mocks base method.
func (m *MockRewardsService) GetRewards(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool) []*vega.Reward {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewards", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*vega.Reward)
	return ret0
}

// GetRewards indicates an expected call of GetRewards.
func (mr *MockRewardsServiceMockRecorder) GetRewards(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewards", reflect.TypeOf((*MockRewardsService)(nil).GetRewards), arg0, arg1, arg2, arg3, arg4)
}

// GetRewardsForAsset mocks base method.
func (m *MockRewardsService) GetRewardsForAsset(arg0 context.Context, arg1, arg2 string, arg3, arg4 uint64, arg5 bool) []*vega.Reward {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardsForAsset", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*vega.Reward)
	return ret0
}

// GetRewardsForAsset indicates an expected call of GetRewardsForAsset.
func (mr *MockRewardsServiceMockRecorder) GetRewardsForAsset(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardsForAsset", reflect.TypeOf((*MockRewardsService)(nil).GetRewardsForAsset), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ObserveRewards mocks base method.
func (m *MockRewardsService) ObserveRewards(arg0 context.Context, arg1 int, arg2, arg3 string) (<-chan vega.Reward, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveRewards", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan vega.Reward)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveRewards indicates an expected call of ObserveRewards.
func (mr *MockRewardsServiceMockRecorder) ObserveRewards(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveRewards", reflect.TypeOf((*MockRewardsService)(nil).ObserveRewards), arg0, arg1, arg2, arg3)
}
