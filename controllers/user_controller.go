package controllers

import (
	"github.com/cosmos-sajal/go_boilerplate/models"
	"github.com/cosmos-sajal/go_boilerplate/validators"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var requestBody validators.CreateUser

	c.Bind(&requestBody)
	validationErr := requestBody.Validate()
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	user, err := models.CreateUser(
		*requestBody.Name, *requestBody.MobileNumber, *requestBody.Dob)
	if err != nil {
		c.JSON(400, err.Error())
	}

	c.JSON(200, user)
}
