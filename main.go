package main

import (
	"github.com/cosmos-sajal/go_boilerplate/controllers"
	"github.com/cosmos-sajal/go_boilerplate/initializers"
	authservice "github.com/cosmos-sajal/go_boilerplate/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/user/signin/", controllers.SignInController)
	r.POST("/api/v1/otp/validate/", controllers.OTPValidateController)
	r.POST("/api/v1/token/refresh/", controllers.RefreshTokenController)

	authGroup := r.Group("")
	authGroup.Use(authservice.JWTAuthMiddleware())
	authGroup.POST("/api/v1/user/", controllers.CreateUser)
	authGroup.GET("/api/v1/user/", controllers.GetUserList)
	authGroup.PATCH("/api/v1/user/:user_id/", controllers.UpdateUser)

	r.Run() // listen and serve on 0.0.0.0:3000
}
