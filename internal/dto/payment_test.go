package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentCheckout_JSONMarshalling(t *testing.T) {
	meta := MetaData{}
	payment := PaymentCheckout{
		OrderId:     "order-1",
		TotalAmount: 99.99,
		Products: []PaymentItemCheckoutProduct{
			{ProductId: "prod-1", Quantity: 2},
			{ProductId: "prod-2", Quantity: 1},
		},
		MetaData: meta,
	}
	data, err := json.Marshal(payment)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"orderId":"order-1"`)
	assert.Contains(t, string(data), `"totalAmount":99.99`)
	assert.Contains(t, string(data), `"productId":"prod-1"`)
	assert.Contains(t, string(data), `"quantity":2`)
}

func TestPaymentItemCheckoutProduct_Fields(t *testing.T) {
	item := PaymentItemCheckoutProduct{
		ProductId: "prod-3",
		Quantity:  5,
	}
	assert.Equal(t, "prod-3", item.ProductId)
	assert.Equal(t, 5, item.Quantity)
}
