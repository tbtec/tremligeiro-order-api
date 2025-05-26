package dto

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder_JSONMarshalling(t *testing.T) {
	c := CreateOrder{CustomerId: "c1"}
	data, err := json.Marshal(c)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"customerId":"c1"`)
}

func TestOrder_JSONMarshalling(t *testing.T) {
	now := time.Now()
	customerId := "c2"
	order := Order{
		ID:          "o1",
		CustomerId:  &customerId,
		Status:      "PENDING",
		TotalAmount: 123.45,
		CreatedAt:   now,
		UpdatedAt:   now,
		MetaData:    MetaData{},
	}
	data, err := json.Marshal(order)
	assert.NoError(t, err)

	var unmarshalled Order
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, unmarshalled.ID)
	assert.NotNil(t, unmarshalled.CustomerId)
	assert.Equal(t, *order.CustomerId, *unmarshalled.CustomerId)
	assert.Equal(t, order.Status, unmarshalled.Status)
	assert.Equal(t, order.TotalAmount, unmarshalled.TotalAmount)
	assert.WithinDuration(t, order.CreatedAt, unmarshalled.CreatedAt, time.Second)
	assert.WithinDuration(t, order.UpdatedAt, unmarshalled.UpdatedAt, time.Second)
}

func TestOrderCheckout_JSONMarshalling(t *testing.T) {
	meta := MetaData{}
	oc := OrderCheckout{
		OrderId: "o2",
		Products: []OrderCheckoutProduct{
			{ProductId: "p1", Quantity: 2},
			{ProductId: "p2", Quantity: 1},
		},
		MetaData: meta,
	}
	data, err := json.Marshal(oc)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"orderId":"o2"`)
	assert.Contains(t, string(data), `"productId":"p1"`)
	assert.Contains(t, string(data), `"quantity":2`)
}

func TestOrderContent_JSONMarshalling(t *testing.T) {
	now := time.Now()
	customerId := "c3"
	content := OrderContent{
		Content: []Order{
			{
				ID:          "o3",
				CustomerId:  &customerId,
				Status:      "READY",
				TotalAmount: 50.0,
				CreatedAt:   now,
				UpdatedAt:   now,
				MetaData:    MetaData{},
			},
		},
	}
	data, err := json.Marshal(content)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"o3"`)
	assert.Contains(t, string(data), `"customerId":"c3"`)
}

func TestOrderDetails_JSONMarshalling(t *testing.T) {
	now := time.Now()
	customerId := "c4"
	details := OrderDetails{
		ID:          "o4",
		CustomerId:  &customerId,
		Status:      "FINALIZED",
		TotalAmount: 200.0,
		CreatedAt:   now,
		UpdatedAt:   now,
		OrderProducts: []OrderProduct{
			{ID: "op1", ProductID: "p1", Quantity: 1},
		},
	}
	data, err := json.Marshal(details)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"o4"`)
	assert.Contains(t, string(data), `"customerId":"c4"`)
	assert.Contains(t, string(data), `"orderProducts"`)
}

func TestUpdateOrder_JSONMarshalling(t *testing.T) {
	upd := UpdateOrder{Status: "READY"}
	data, err := json.Marshal(upd)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"status":"READY"`)
}

func TestOrderEvent_JSONMarshalling(t *testing.T) {
	evt := OrderEvent{ID: "o5", Status: "PENDING"}
	data, err := json.Marshal(evt)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"o5"`)
	assert.Contains(t, string(data), `"status":"PENDING"`)
}
