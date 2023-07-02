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

func OTPValidateController(c *gin.Context) {
	var requestBody validators.OTPValidatorStruct
	c.Bind(&requestBody)
	validationErr := requestBody.Validate()
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	user, err := models.GetUserByMobile(*requestBody.MobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	tokenStruct, err := authservice.GenerateToken(int(user.ID))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	helpers.DeleteKey("OTP_KEY_" + user.MobileNumber)

	c.JSON(http.StatusOK, tokenStruct)
}
