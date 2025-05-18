// internal/core/gateway/product_test.go
package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

// Mock for IProductRepository (not used in FindOne, but required for constructor)
type mockProductRepository struct {
	repository.IProductRepository
}

// Mock for IProductService
type mockProductService struct {
	mock.Mock
	external.IProductService
}

func (m *mockProductService) FindOne(ctx context.Context, id string) (*external.ProductResponse, error) {
	args := m.Called(ctx, id)
	if pr, ok := args.Get(0).(*external.ProductResponse); ok {
		return pr, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestProductGateway_FindOne_Success(t *testing.T) {
	mockRepo := &mockProductRepository{}
	mockService := &mockProductService{}

	now := time.Now()
	productResp := &external.ProductResponse{
		ProductId:   "p1",
		Name:        "Test Product",
		Description: "Desc",
		Amount:      10.5,
		Category:    external.CategoryResponse{ID: 1},
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockService.On("FindOne", mock.Anything, "p1").Return(productResp, nil)

	gateway := NewProductGateway(mockRepo, mockService)
	product, err := gateway.FindOne(context.Background(), "p1")

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "p1", product.ID)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, "Desc", product.Description)
	assert.Equal(t, 10.5, product.Amount)
	assert.Equal(t, 1, product.CategoryId)
	assert.WithinDuration(t, now, product.CreatedAt, time.Second)
	assert.WithinDuration(t, now, product.UpdatedAt, time.Second)
}

func TestProductGateway_FindOne_NotFound(t *testing.T) {
	mockRepo := &mockProductRepository{}
	mockService := &mockProductService{}

	mockService.On("FindOne", mock.Anything, "notfound").Return(nil, errors.New("not found"))

	gateway := NewProductGateway(mockRepo, mockService)
	product, err := gateway.FindOne(context.Background(), "notfound")

	assert.Nil(t, product)
	assert.EqualError(t, err, "not found")
}
