package authservice

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateToken(userId int) (*Token, error) {
	secretKey := []byte(os.Getenv("JWT_TOKEN"))
	// Generating Access Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	})
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	// Generating Refresh TOken
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"sub":     1,
	})
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, err
}

func IsRateLimitExceeded(mobileNumber string) bool {
	key := "OTP_ATTEMPTS_" + mobileNumber
	val, err := helpers.GetCacheValue(key)
	if err != nil {
		return false
	}
	num, _ := strconv.Atoi(val)
	fmt.Println(key, num, val)

	return num > 5
}

func IncrementOTPAttemptCounter(mobileNumber string) {
	key := "OTP_ATTEMPTS_" + mobileNumber
	val, err := helpers.GetCacheValue(key)
	if err != nil {
		helpers.SetCacheValue(key, "1", 3600)
		return
	}
	num, _ := strconv.Atoi(val)
	num++
	fmt.Println(key, num, val)
	helpers.SetCacheValue(key, strconv.Itoa(num), 3600)
}
