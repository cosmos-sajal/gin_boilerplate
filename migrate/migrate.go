package main

import (
	"github.com/cosmos-sajal/go_boilerplate/initializers"
	"github.com/cosmos-sajal/go_boilerplate/models"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	db := initializers.ConnectToDB()
	db.AutoMigrate(&models.User{})
}
