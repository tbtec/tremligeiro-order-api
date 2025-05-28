package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCmdCreateProduct_ToNewEntity(t *testing.T) {
	cmd := CmdCreateProduct{
		Name:        "Coca-Cola",
		Description: "Refrigerante",
		CategoryId:  1,
		Amount:      5.50,
	}
	prod := cmd.ToNewEntity()
	assert.NotEmpty(t, prod.ID)
	assert.Equal(t, "Coca-Cola", prod.Name)
	assert.Equal(t, "Refrigerante", prod.Description)
	assert.Equal(t, 1, prod.CategoryId)
	assert.Equal(t, 5.50, prod.Amount)
	assert.WithinDuration(t, time.Now().UTC(), prod.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now().UTC(), prod.UpdatedAt, time.Second)
}

func TestCmdUpdateProduct_ToUpdateEntity(t *testing.T) {
	now := time.Now()
	cmd := CmdUpdateProduct{
		ProductId:   "p1",
		Name:        "Pepsi",
		Description: "Refrigerante de cola",
		CategoryId:  2,
		Amount:      4.99,
		CreatedAt:   now,
	}
	prod := cmd.ToUpdateEntity()
	assert.Equal(t, "p1", prod.ID)
	assert.Equal(t, "Pepsi", prod.Name)
	assert.Equal(t, "Refrigerante de cola", prod.Description)
	assert.Equal(t, 2, prod.CategoryId)
	assert.Equal(t, 4.99, prod.Amount)
	assert.Equal(t, now, prod.CreatedAt)
	assert.WithinDuration(t, time.Now().UTC(), prod.UpdatedAt, time.Second)
}

func TestCreateProductOutput_Fields(t *testing.T) {
	now := time.Now()
	out := CreateProductOutput{
		ProductId:    "p2",
		Name:         "Guaraná",
		Description:  "Refrigerante de guaraná",
		Amount:       3.75,
		CategoryID:   3,
		CategoryName: "Bebidas",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	assert.Equal(t, "p2", out.ProductId)
	assert.Equal(t, "Guaraná", out.Name)
	assert.Equal(t, "Refrigerante de guaraná", out.Description)
	assert.Equal(t, 3, out.CategoryID)
	assert.Equal(t, "Bebidas", out.CategoryName)
	assert.Equal(t, 3.75, out.Amount)
	assert.Equal(t, now, out.CreatedAt)
	assert.Equal(t, now, out.UpdatedAt)
}
