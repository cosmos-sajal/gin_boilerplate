package queueservice

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var QueueService *sqs.SQS

func pollQueues() {
	smsSenderChnMessages := make(chan *sqs.Message)
	go PollSMSQueue(smsSenderChnMessages)
	for message := range smsSenderChnMessages {
		stringifiedMessage := *message.Body
		var sendSMSStruct SendSMSStruct
		err := json.Unmarshal([]byte(stringifiedMessage), &sendSMSStruct)
		if err != nil {
			log.Fatal("Failed to convert JSON to struct:", err)
		}

		sendSMSStruct.ConsumeMessage()
		sendSMSStruct.DeleteMessage(message)
	}
}

func ConnectToAsyncClient() (*sqs.SQS, error) {
	sqsRegion := os.Getenv("SQS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(sqsRegion),
	})
	if err != nil {
		return nil, err
	}
	QueueService = sqs.New(sess)
	pollQueues()

	return QueueService, nil
}
