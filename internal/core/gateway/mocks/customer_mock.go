// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/gateway/customer.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/gateway/customer.go -destination=internal/core/gateway/mocks/customer_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/tbtec/tremligeiro/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockICustomerGateway is a mock of ICustomerGateway interface.
type MockICustomerGateway struct {
	ctrl     *gomock.Controller
	recorder *MockICustomerGatewayMockRecorder
	isgomock struct{}
}

// MockICustomerGatewayMockRecorder is the mock recorder for MockICustomerGateway.
type MockICustomerGatewayMockRecorder struct {
	mock *MockICustomerGateway
}

// NewMockICustomerGateway creates a new mock instance.
func NewMockICustomerGateway(ctrl *gomock.Controller) *MockICustomerGateway {
	mock := &MockICustomerGateway{ctrl: ctrl}
	mock.recorder = &MockICustomerGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICustomerGateway) EXPECT() *MockICustomerGatewayMockRecorder {
	return m.recorder
}

// FindByDocumentNumber mocks base method.
func (m *MockICustomerGateway) FindByDocumentNumber(ctx context.Context, documentNumber string) (*entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDocumentNumber", ctx, documentNumber)
	ret0, _ := ret[0].(*entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDocumentNumber indicates an expected call of FindByDocumentNumber.
func (mr *MockICustomerGatewayMockRecorder) FindByDocumentNumber(ctx, documentNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDocumentNumber", reflect.TypeOf((*MockICustomerGateway)(nil).FindByDocumentNumber), ctx, documentNumber)
}

// FindOne mocks base method.
func (m *MockICustomerGateway) FindOne(ctx context.Context, id string) (*entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, id)
	ret0, _ := ret[0].(*entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockICustomerGatewayMockRecorder) FindOne(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockICustomerGateway)(nil).FindOne), ctx, id)
}
