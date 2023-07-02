package main

import (
	"github.com/cosmos-sajal/go_boilerplate/controllers"
	"github.com/cosmos-sajal/go_boilerplate/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/api/v1/user/", controllers.CreateUser)
	r.GET("/api/v1/user/", controllers.GetUserList)
	r.Run() // listen and serve on 0.0.0.0:8080
}
