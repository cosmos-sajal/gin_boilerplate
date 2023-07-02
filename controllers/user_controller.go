package controllers

import (
	"log"

	"github.com/cosmos-sajal/go_boilerplate/models"
	"github.com/cosmos-sajal/go_boilerplate/validators"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

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

func GetUserList(c *gin.Context) {
	var queryParams validators.GetUserList
	err := decoder.Decode(&queryParams, c.Request.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	}

	validationErr := queryParams.Validate()
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	userList, err := models.GetUserList(*queryParams.Limit, *queryParams.Offset)
	if err != nil {
		c.JSON(400, err.Error())
	}

	c.JSON(200, userList)
}
