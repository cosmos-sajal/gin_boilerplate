package queueservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	SMS_SENDER = "local-gin-boilerplate-send-sms"
)

type SQSInterface interface {
	SendMessage() (*sqs.SendMessageOutput, error)
	ConsumeMessage() error
	DeleteMessage(*sqs.Message) error
}

func getQueueURL(queueName string) *string {
	queueURL, err := QueueService.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return queueURL.QueueUrl
	}
}
