package main

import (
	"os"
	"strconv"

	"github.com/cosmos-sajal/go_boilerplate/controllers"
	"github.com/cosmos-sajal/go_boilerplate/initializers"
	logger "github.com/cosmos-sajal/go_boilerplate/logger_middeware"
	authservice "github.com/cosmos-sajal/go_boilerplate/services/auth"
	"github.com/cosmos-sajal/go_boilerplate/worker"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
	initializers.ConnectToAsyncClient()
}

func main() {
	isAppServer, _ := strconv.ParseBool(os.Getenv("IS_APP_SERVER"))
	if isAppServer {
		r := gin.Default()
		r.Use(logger.RequestResponseLoggerMiddleware())

		r.POST("/api/v1/user/signin/", controllers.SignInController)
		r.POST("/api/v1/otp/validate/", controllers.OTPValidateController)
		r.POST("/api/v1/token/refresh/", controllers.RefreshTokenController)

		authGroup := r.Group("")
		authGroup.Use(authservice.JWTAuthMiddleware())
		authGroup.POST("/api/v1/user/", controllers.CreateUser)
		authGroup.GET("/api/v1/user/", controllers.GetUserList)
		authGroup.PATCH("/api/v1/user/:user_id/", controllers.UpdateUser)

		r.Run() // listen and serve on 0.0.0.0:3000
	} else {
		worker.StartWorker(initializers.TaskServer, os.Getenv("QUEUE_NAME"))
		initializers.InitialiseCron()
	}
}
