package tasks

import (
	"fmt"
)

type EmailStruct struct {
	UserId    int
	EmailBody string
}

func SendEmail(b64payload string) error {
	payload := EmailStruct{}
	DecodeToTask(b64payload, &payload)

	fmt.Println("Consuming Email message", payload.UserId, payload.EmailBody)

	return nil
}
