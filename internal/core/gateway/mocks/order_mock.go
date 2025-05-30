// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/gateway/order.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/gateway/order.go -destination=internal/core/gateway/mocks/order_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/tbtec/tremligeiro/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIOrderGateway is a mock of IOrderGateway interface.
type MockIOrderGateway struct {
	ctrl     *gomock.Controller
	recorder *MockIOrderGatewayMockRecorder
	isgomock struct{}
}

// MockIOrderGatewayMockRecorder is the mock recorder for MockIOrderGateway.
type MockIOrderGatewayMockRecorder struct {
	mock *MockIOrderGateway
}

// NewMockIOrderGateway creates a new mock instance.
func NewMockIOrderGateway(ctrl *gomock.Controller) *MockIOrderGateway {
	mock := &MockIOrderGateway{ctrl: ctrl}
	mock.recorder = &MockIOrderGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOrderGateway) EXPECT() *MockIOrderGatewayMockRecorder {
	return m.recorder
}

// FindOne mocks base method.
func (m *MockIOrderGateway) FindOne(ctx context.Context, id string) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, id)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockIOrderGatewayMockRecorder) FindOne(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIOrderGateway)(nil).FindOne), ctx, id)
}

// Update mocks base method.
func (m *MockIOrderGateway) Update(ctx context.Context, order *entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIOrderGatewayMockRecorder) Update(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIOrderGateway)(nil).Update), ctx, order)
}
