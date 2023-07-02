package controllers

import (
	"log"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusCreated, user)
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
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, userList)
}

func UpdateUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("user_id"))
	var requestBody validators.UpdateUser
	c.Bind(&requestBody)
	validationErr := requestBody.Validate(userId)
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	userStruct := models.UpdateUserStruct{
		Name:         requestBody.Name,
		MobileNumber: requestBody.MobileNumber,
		DOB:          requestBody.Dob,
		IsDeleted:    requestBody.IsDeleted,
	}
	user, err := models.UpdateUser(userId, &userStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
