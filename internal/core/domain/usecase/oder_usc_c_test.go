package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"

	"github.com/tbtec/tremligeiro/internal/core/gateway/mocks"

	"github.com/tbtec/tremligeiro/internal/dto"
	"go.uber.org/mock/gomock"
)

func TestCreateOrder_WithCustomer_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderGateway := mocks.NewMockIOrderGateway(ctrl)
	mockCustomerGateway := mocks.NewMockICustomerGateway(ctrl)

	usecase := usecase.NewUseCaseCreateOrder(mockOrderGateway, mockCustomerGateway)

	createOrder := dto.CreateOrder{
		CustomerId: "cust-123",
		// outros campos necessários...
	}

	mockCustomerGateway.EXPECT().FindOne(gomock.Any(), "cust-123").Return(&entity.Customer{}, nil)

	result, err := usecase.Create(context.Background(), createOrder)

	assert.NoError(t, err)
	assert.Equal(t, createOrder.CustomerId, result.CustomerId)
}
