package dto

import (
    "encoding/json"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestMetaData_JSONMarshalling(t *testing.T) {
    paymentId := "pay-123"
    webhookUrl := "http://webhook.url"
    meta := MetaData{
        PaymentId:         &paymentId,
        PaymentWebHookUrl: &webhookUrl,
    }
    data, err := json.Marshal(meta)
    assert.NoError(t, err)
    assert.Contains(t, string(data), `"paymentId":"pay-123"`)
    assert.Contains(t, string(data), `"paymentWebhookUrl":"http://webhook.url"`)

    var unmarshalled MetaData
    err = json.Unmarshal(data, &unmarshalled)
    assert.NoError(t, err)
    assert.NotNil(t, unmarshalled.PaymentId)
    assert.NotNil(t, unmarshalled.PaymentWebHookUrl)
    assert.Equal(t, paymentId, *unmarshalled.PaymentId)
    assert.Equal(t, webhookUrl, *unmarshalled.PaymentWebHookUrl)
}

func TestMetaDataContent_JSONMarshalling(t *testing.T) {
    paymentId := "pay-456"
    webhookUrl := "http://webhook2.url"
    meta := MetaData{
        PaymentId:         &paymentId,
        PaymentWebHookUrl: &webhookUrl,
    }
    content := MetaDataContent{Content: meta}
    data, err := json.Marshal(content)
    assert.NoError(t, err)
    assert.Contains(t, string(data), `"paymentId":"pay-456"`)
    assert.Contains(t, string(data), `"paymentWebhookUrl":"http://webhook2.url"`)

    var unmarshalled MetaDataContent
    err = json.Unmarshal(data, &unmarshalled)
    assert.NoError(t, err)
    assert.NotNil(t, unmarshalled.Content.PaymentId)
    assert.NotNil(t, unmarshalled.Content.PaymentWebHookUrl)
    assert.Equal(t, paymentId, *unmarshalled.Content.PaymentId)
    assert.Equal(t, webhookUrl, *unmarshalled.Content.PaymentWebHookUrl)
}