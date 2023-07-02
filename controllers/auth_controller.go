package controllers

import (
	"fmt"
	"net/http"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/models"
	"github.com/cosmos-sajal/go_boilerplate/validators"
	"github.com/gin-gonic/gin"
)

func SignInController(c *gin.Context) {
	var requestBody validators.SignInStruct
	var randomOTP string
	c.Bind(&requestBody)
	validationErr := requestBody.Validate()
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	user, _ := models.GetUserByMobile(*requestBody.MobileNumber)
	val, _ := helpers.GetCacheValue("OTP_KEY_" + user.MobileNumber)
	if val == "" {
		randomOTP = helpers.GenerateRandomOTP()
		helpers.SetCacheValue("OTP_KEY_"+user.MobileNumber, randomOTP, 120)
	} else {
		randomOTP = val
	}

	fmt.Println("OTP is - " + randomOTP)

	// helpers.SendOTP(*requestBody.MobileNumber, randomOTP)
	c.JSON(http.StatusOK, gin.H{
		"result": "OTP Sent successfully",
	})
}
