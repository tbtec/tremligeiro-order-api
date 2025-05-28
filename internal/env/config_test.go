package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvConfigDefaults(t *testing.T) {
	// Remove variáveis de ambiente para testar defaults
	os.Unsetenv("ENV")
	os.Unsetenv("PORT")

	cfg, err := LoadEnvConfig()
	assert.NoError(t, err)
	assert.Equal(t, 8083, cfg.Port)
}

func TestLoadEnvConfigWithEnvVars(t *testing.T) {
	os.Setenv("ENV", "test")
	os.Setenv("PORT", "1234")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_USER", "user")
	os.Setenv("POSTGRES_PASS", "pass")
	os.Setenv("POSTGRES_DB", "db")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("PAYMENT_URL", "http://pay")
	os.Setenv("PAYMENT_AUTH_TOKEN", "token")
	os.Setenv("CUSTOMER_URL", "http://cust")
	os.Setenv("PRODUCT_URL", "http://prod")
	os.Setenv("ORDER_TOPIC_ARN", "arn")
	os.Setenv("PRODUCTION_ORDER_QUEUE_URL", "queue")
	os.Setenv("AWS_ACCESS_KEY_ID", "key")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_SESSION_TOKEN", "token")
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_USE_CREDENTIALS", "true")

	cfg, err := LoadEnvConfig()
	assert.NoError(t, err)
	assert.Equal(t, "test", cfg.Env)
	assert.Equal(t, 1234, cfg.Port)
	assert.Equal(t, "localhost", cfg.DbHost)
	assert.Equal(t, "user", cfg.DbUser)
	assert.Equal(t, "pass", cfg.DbPassword)
	assert.Equal(t, "db", cfg.DbName)
	assert.Equal(t, 5432, cfg.DbPort)
	assert.Equal(t, "http://pay", cfg.PaymentUrl)
	assert.Equal(t, "token", cfg.PaymentAuthToken)
	assert.Equal(t, "http://cust", cfg.CustomerUrl)
	assert.Equal(t, "http://prod", cfg.ProductUrl)
	assert.Equal(t, "arn", cfg.OrderTopicArn)
	assert.Equal(t, "queue", cfg.ProductionOrderQueueUrl)
	assert.Equal(t, "key", cfg.AwsAccessKeyId)
	assert.Equal(t, "secret", cfg.AwsSecretAccessKey)
	assert.Equal(t, "token", cfg.AwsSessionToken)
	assert.Equal(t, "us-east-1", cfg.AwsRegion)
	assert.Equal(t, "true", cfg.AwsUseCredentials)
}
