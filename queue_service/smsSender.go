package queueservice

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/cosmos-sajal/go_boilerplate/communication"
)

type SendSMSStruct struct {
	OTP          string
	MobileNumber string
	UserId       int
}

func (s *SendSMSStruct) SendMessage() (*sqs.SendMessageOutput, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Failed to convert struct to JSON:", err)
	}
	jsonString := string(jsonData)

	return QueueService.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(jsonString),
		QueueUrl:    getQueueURL(SMS_SENDER),
	})
}

func (s *SendSMSStruct) ConsumeMessage() error {
	fmt.Println("Consuming OTP message")
	otpBody := communication.OTPBodyStruct{
		OTP: s.OTP,
		To:  s.MobileNumber,
	}

	return otpBody.SendSMS()
}

func (s *SendSMSStruct) DeleteMessage(message *sqs.Message) error {
	QueueService.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      getQueueURL(SMS_SENDER),
		ReceiptHandle: message.ReceiptHandle,
	})

	return nil
}

func PollSMSQueue(chn chan<- *sqs.Message) {
	for {
		output, err := QueueService.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: getQueueURL(SMS_SENDER),
		})
		if err != nil {
			fmt.Printf("failed to fetch sqs message %v", err)
		}

		for _, message := range output.Messages {
			chn <- message
		}
	}
}
