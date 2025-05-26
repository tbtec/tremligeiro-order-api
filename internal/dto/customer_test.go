package dto

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer_JSONTags(t *testing.T) {
	c := CreateCustomer{
		Name:           "João",
		DocumentNumber: "12345678900",
		Email:          "joao@email.com",
	}
	data, err := json.Marshal(c)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"name":"João"`)
	assert.Contains(t, string(data), `"documentNumber":"12345678900"`)
	assert.Contains(t, string(data), `"email":"joao@email.com"`)
}

func TestUpdateCustomer_Fields(t *testing.T) {
	now := time.Now()
	u := UpdateCustomer{
		CustomerId:     "c1",
		Name:           "Maria",
		DocumentNumber: "98765432100",
		Email:          "maria@email.com",
		UpdatedAt:      now,
	}
	assert.Equal(t, "c1", u.CustomerId)
	assert.Equal(t, "Maria", u.Name)
	assert.Equal(t, "98765432100", u.DocumentNumber)
	assert.Equal(t, "maria@email.com", u.Email)
	assert.Equal(t, now, u.UpdatedAt)
}

func TestFindCustomer_JSONTags(t *testing.T) {
	f := FindCustomer{DocumentNumber: "11122233344"}
	data, err := json.Marshal(f)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"documentNumber":"11122233344"`)
}

func TestCustomer_JSONMarshalling(t *testing.T) {
	now := time.Now()
	c := Customer{
		CustomerId:     "c2",
		Name:           "Pedro",
		DocumentNumber: "55566677788",
		Email:          "pedro@email.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshalled Customer
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}

func TestCustomerContent_JSONMarshalling(t *testing.T) {
	now := time.Now()
	c := Customer{
		CustomerId:     "c3",
		Name:           "Ana",
		DocumentNumber: "99988877766",
		Email:          "ana@email.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	content := CustomerContent{Content: c}
	data, err := json.Marshal(content)
	assert.NoError(t, err)

	var unmarshalled CustomerContent
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}
