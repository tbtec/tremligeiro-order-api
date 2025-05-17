package event

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type IProducerService interface {
	PublishMessage(ctx context.Context, order dto.Order) error
}

type ProducerService struct {
	TopicName string
	TopicArn  string
	Client    *sns.Client
}

func NewProducerService(topicArn string, config aws.Config) IProducerService {
	return &ProducerService{
		TopicArn: topicArn,
		Client:   sns.NewFromConfig(config),
	}
}
func (producer *ProducerService) PublishMessage(ctx context.Context, order dto.Order) error {
	// Serialize order to JSON
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, err = producer.Client.Publish(ctx, &sns.PublishInput{
		Message:  aws.String(string(orderBytes)),
		TopicArn: aws.String(producer.TopicArn),
	})

	if err != nil {
		return err
	}

	return nil
}
