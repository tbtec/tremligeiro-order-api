package dto

import (
    "encoding/json"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCategory_JSONMarshalling(t *testing.T) {
    cat := Category{
        ID:   1,
        Name: "Bebidas",
    }

    data, err := json.Marshal(cat)
    assert.NoError(t, err)

    var unmarshalled Category
    err = json.Unmarshal(data, &unmarshalled)
    assert.NoError(t, err)
    assert.Equal(t, cat, unmarshalled)
}

func TestCategory_Fields(t *testing.T) {
    cat := Category{ID: 10, Name: "Alimentos"}
    assert.Equal(t, 10, cat.ID)
    assert.Equal(t, "Alimentos", cat.Name)
}