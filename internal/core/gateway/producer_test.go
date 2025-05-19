package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
)

// Mock para event.IProducerService
type MockProducerService struct {
	mock.Mock
}

func (m *MockProducerService) PublishMessage(ctx context.Context, message interface{}) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}

func TestOrderProducerGatewayPublishMessageSuccess(t *testing.T) {
	mockProducer := new(MockProducerService)
	gateway := NewOrderProducerGateway(mockProducer)

	customerId := "customer-1"
	order := &entity.Order{
		ID:          "order-1",
		CustomerId:  &customerId,
		Status:      "PENDING",
		TotalAmount: 100.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockProducer.On("PublishMessage", mock.Anything, mock.Anything).Return(nil)

	err := gateway.PublishMessage(context.Background(), order)
	assert.NoError(t, err)
	mockProducer.AssertExpectations(t)
}

func TestOrderProducerGatewayPublishMessageError(t *testing.T) {
	mockProducer := new(MockProducerService)
	gateway := NewOrderProducerGateway(mockProducer)
	customerId := "customer-2"
	order := &entity.Order{
		ID:          "order-2",
		CustomerId:  &customerId,
		Status:      "FAILED",
		TotalAmount: 50.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockProducer.On("PublishMessage", mock.Anything, mock.Anything).Return(errors.New("publish error"))

	err := gateway.PublishMessage(context.Background(), order)
	assert.Error(t, err)
	assert.EqualError(t, err, "publish error")
	mockProducer.AssertExpectations(t)
}
