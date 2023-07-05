package initializers

import (
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/cosmos-sajal/go_boilerplate/tasks"
)

var TaskServer *machinery.Server

func ConnectToAsyncClient() {
	var err error
	var cnf = &config.Config{
		Broker:        "redis://redis:6379",
		ResultBackend: "redis://redis:6379",
	}

	TaskServer, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal("Error connecting to Redis", err)
		return
	}

	TaskServer.RegisterTasks(map[string]interface{}{
		"send_otp": tasks.SendOTP,
	})
}
