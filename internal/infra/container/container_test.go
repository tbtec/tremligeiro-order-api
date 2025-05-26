package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/env"
)

func TestNewContainer(t *testing.T) {
	cfg := env.Config{
		Env:                     "test",
		Port:                    8080,
		DbHost:                  "localhost",
		DbUser:                  "user",
		DbPassword:              "pass",
		DbName:                  "db",
		DbPort:                  5432,
		PaymentUrl:              "http://pay",
		PaymentAuthToken:        "token",
		CustomerUrl:             "http://cust",
		ProductUrl:              "http://prod",
		OrderTopicArn:           "arn",
		ProductionOrderQueueUrl: "queue",
		AwsAccessKeyId:          "key",
		AwsSecretAccessKey:      "secret",
		AwsSessionToken:         "token",
		AwsRegion:               "us-east-1",
		AwsUseCredentials:       "true",
	}
	c, err := New(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, cfg, c.Config)
}

func TestGetPostgreSQLConf(t *testing.T) {
	cfg := env.Config{
		DbUser:     "user",
		DbPassword: "pass",
		DbHost:     "localhost",
		DbPort:     5432,
		DbName:     "db",
	}
	conf := getPostgreSQLConf(cfg)
	assert.Equal(t, "user", conf.User)
	assert.Equal(t, "pass", conf.Pass)
	assert.Equal(t, "localhost", conf.Url)
	assert.Equal(t, 5432, conf.Port)
	assert.Equal(t, "db", conf.DbName)
}

func TestGetPaymentConf(t *testing.T) {
	cfg := env.Config{
		PaymentUrl:       "http://pay",
		PaymentAuthToken: "token",
	}
	conf := getPaymentConf(cfg)
	assert.Equal(t, "http://pay", conf.Url)
	assert.Equal(t, "token", conf.Token)
}

func TestGetCustomerConf(t *testing.T) {
	cfg := env.Config{
		CustomerUrl: "http://cust",
	}
	conf := getCustomerConf(cfg)
	assert.Equal(t, "http://cust", conf.Url)
}

func TestGetProductConf(t *testing.T) {
	cfg := env.Config{
		ProductUrl: "http://prod",
	}
	conf := getProductConf(cfg)
	assert.Equal(t, "http://prod", conf.Url)
}
