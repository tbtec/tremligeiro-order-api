package dto

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_JSONMarshalling(t *testing.T) {
	cp := CreateProduct{
		Name:        "Coca-Cola",
		Description: "Refrigerante",
		CategoryId:  1,
		Amount:      5.50,
	}
	data, err := json.Marshal(cp)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"name":"Coca-Cola"`)
	assert.Contains(t, string(data), `"description":"Refrigerante"`)
	assert.Contains(t, string(data), `"categoryId":1`)
	assert.Contains(t, string(data), `"amount":5.5`)
}

func TestUpdateProduct_Fields(t *testing.T) {
	now := time.Now()
	up := UpdateProduct{
		ProductId:   "p1",
		Name:        "Pepsi",
		Description: "Refrigerante",
		CategoryId:  2,
		Amount:      4.99,
		CreatedAt:   now,
	}
	assert.Equal(t, "p1", up.ProductId)
	assert.Equal(t, "Pepsi", up.Name)
	assert.Equal(t, "Refrigerante", up.Description)
	assert.Equal(t, 2, up.CategoryId)
	assert.Equal(t, 4.99, up.Amount)
	assert.Equal(t, now, up.CreatedAt)
}

func TestProduct_JSONMarshalling(t *testing.T) {
	now := time.Now()
	prod := Product{
		ProductId:   "p2",
		Name:        "Guaraná",
		Description: "Refrigerante de guaraná",
		Amount:      3.75,
		Category:    Category{ID: 3, Name: "Bebidas"},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	data, err := json.Marshal(prod)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"p2"`)
	assert.Contains(t, string(data), `"name":"Guaraná"`)
	assert.Contains(t, string(data), `"description":"Refrigerante de guaraná"`)
	assert.Contains(t, string(data), `"amount":3.75`)
	assert.Contains(t, string(data), `"category":{"id":3,"name":"Bebidas"}`)
}

func TestProductContent_JSONMarshalling(t *testing.T) {
	now := time.Now()
	products := []Product{
		{
			ProductId:   "p3",
			Name:        "Água",
			Description: "Água mineral",
			Amount:      2.00,
			Category:    Category{ID: 1, Name: "Bebidas"},
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	content := ProductContent{Content: products}
	data, err := json.Marshal(content)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"p3"`)
	assert.Contains(t, string(data), `"name":"Água"`)
	assert.Contains(t, string(data), `"category":{"id":1,"name":"Bebidas"}`)
}
