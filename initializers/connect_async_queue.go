package initializers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/cosmos-sajal/go_boilerplate/tasks"
	"github.com/cosmos-sajal/go_boilerplate/worker"
)

var TaskServer *machinery.Server

func isProd() bool {
	return os.Getenv("ENV") == "prod"
}

func ConnectToAsyncClient() {
	var err error
	var cnf *config.Config

	if !isProd() {
		cnf = &config.Config{
			Broker:        os.Getenv("BROKER_BACKEND"),
			ResultBackend: os.Getenv("BROKER_RESULT_BACKEND"),
		}
	} else {
		var sqsClient = sqs.New(session.Must(session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("SQS_REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
			HTTPClient: &http.Client{
				Timeout: time.Second * 120,
			},
		})))
		var visibilityTimeout = 20
		cnf = &config.Config{
			Broker:        os.Getenv("SQS_URL"),
			DefaultQueue:  "local_machinery_tasks",
			ResultBackend: os.Getenv("BROKER_RESULT_BACKEND"),
			SQS: &config.SQSConfig{
				Client: sqsClient,
				// if VisibilityTimeout is nil default to the overall visibility timeout setting for the queue
				// https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html
				VisibilityTimeout: &visibilityTimeout,
				WaitTimeSeconds:   20,
			},
		}
	}

	TaskServer, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal("Error connecting to Redis", err)
		return
	}

	TaskServer.RegisterTasks(map[string]interface{}{
		worker.SEND_OTP_QUEUE:   tasks.SendOTP,
		worker.SEND_EMAIL_QUEUE: tasks.SendEmail,
	})
}
