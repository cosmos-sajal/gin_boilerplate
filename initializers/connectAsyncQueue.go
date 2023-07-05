package initializers

import (
	"log"
	"os"
	"strconv"

	queueservice "github.com/cosmos-sajal/go_boilerplate/queue_service"
)

func ConnectToAsyncClient() {
	shouldUseQueueingSystem, err :=
		strconv.ParseBool(os.Getenv("SHOULD_USE_QUEUEING_SYSTEM"))
	if err != nil {
		return
	}
	if !shouldUseQueueingSystem {
		log.Println("SHOULD_USE_QUEUEING_SYSTEM is false")
		return
	}

	_, err = queueservice.ConnectToAsyncClient()
	if err != nil {
		log.Fatal("Error connecting to SQS", err)
	}
}
