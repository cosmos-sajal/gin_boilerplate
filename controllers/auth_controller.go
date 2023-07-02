package controllers

import (
	"fmt"
	"net/http"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/models"
	authservice "github.com/cosmos-sajal/go_boilerplate/services"
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
	cacheKey := authservice.OTP_KEY_PREFIX + user.MobileNumber
	val, _ := helpers.GetCacheValue(cacheKey)
	if val == "" {
		randomOTP = helpers.GenerateRandomOTP()
		helpers.SetCacheValue(cacheKey, randomOTP, 120)
	} else {
		randomOTP = val
	}

	fmt.Println("OTP is - " + randomOTP)

	// helpers.SendOTP(*requestBody.MobileNumber, randomOTP)
	c.JSON(http.StatusOK, gin.H{
		"result": "OTP Sent successfully",
	})
}

func RefreshTokenController(c *gin.Context) {
	var requestBody validators.RefreshTokenStruct
	c.Bind(&requestBody)
	validationErr := requestBody.Validate()
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	userId, _ := authservice.GetUserIdFromToken(*requestBody.RefreshToken)
	tokenStruct, err := authservice.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, tokenStruct)
}
