package worker

import (
	"fmt"
	"os"

	"github.com/RichardKnop/machinery/v1"
)

var env string = os.Getenv("ENV")

var (
	SEND_OTP_QUEUE   = fmt.Sprintf("%s-send-otp", env)
	SEND_EMAIL_QUEUE = fmt.Sprintf("%s-send-email", env)
)

func StartWorker(taskserver *machinery.Server, queueName string) {
	var worker *machinery.Worker

	if env != "prod" {
		worker = taskserver.NewWorker("machinery_worker", 10)
	} else {
		worker = taskserver.NewCustomQueueWorker(queueName, 10, queueName)
	}

	if err := worker.Launch(); err != nil {
		return
	}
}
