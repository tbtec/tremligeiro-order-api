package event

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type IConsumerService interface {
	ConsumeMessage(ctx context.Context) (dto.Order, error)
}

type ConsumerService struct {
	QueueUrl string
	QueueArn string
	Client   *sqs.Client
}

func NewConsumerService(QueueUrl string, config aws.Config) IConsumerService {
	return &ConsumerService{
		QueueUrl: QueueUrl,
		Client:   sqs.NewFromConfig(config),
	}
}

func (consumer *ConsumerService) ConsumeMessage(ctx context.Context) (dto.Order, error) {
	// Receive a message from the queue
	resp, err := consumer.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &consumer.QueueUrl,
		MaxNumberOfMessages: 1,
	})
	if err != nil {
		return dto.Order{}, err
	}

	if len(resp.Messages) == 0 {
		return dto.Order{}, nil // No messages available
	}

	// Deserialize the message body to Order
	var order dto.Order
	err = json.Unmarshal([]byte(*resp.Messages[0].Body), &order)
	if err != nil {
		return dto.Order{}, err
	}

	return order, nil
}

func (consumer *ConsumerService) DeleteMessage(ctx context.Context, receiptHandle string) error {
	_, err := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueUrl,
		ReceiptHandle: &receiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
