package dto

import (
    "encoding/json"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestOrderProduct_JSONMarshalling(t *testing.T) {
    op := OrderProduct{
        ID:        "op1",
        ProductID: "prod1",
        Quantity:  3,
    }
    data, err := json.Marshal(op)
    assert.NoError(t, err)
    assert.Contains(t, string(data), `"id":"op1"`)
    assert.Contains(t, string(data), `"productId":"prod1"`)
    assert.Contains(t, string(data), `"quantity":3`)

    var unmarshalled OrderProduct
    err = json.Unmarshal(data, &unmarshalled)
    assert.NoError(t, err)
    assert.Equal(t, op, unmarshalled)
}

func TestOrderProduct_Fields(t *testing.T) {
    op := OrderProduct{
        ID:        "op2",
        ProductID: "prod2",
        Quantity:  5,
    }
    assert.Equal(t, "op2", op.ID)
    assert.Equal(t, "prod2", op.ProductID)
    assert.Equal(t, int64(5), op.Quantity)
}