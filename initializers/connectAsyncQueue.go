package initializers

import (
	"log"

	queueservice "github.com/cosmos-sajal/go_boilerplate/queue_service"
)

func ConnectToAsyncClient() {
	_, err := queueservice.ConnectToAsyncClient()
	if err != nil {
		log.Fatal("Error connecting to SQS", err)
	}
}
