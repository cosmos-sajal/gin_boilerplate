package authservice

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	ACCESS_TOKEN_TYPE  = "access"
	REFRESH_TOKEN_TYPE = "refresh"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getOTPAttempPrefix(mobileNumber string) string {
	return OTP_ATTEMPTS_PREFIX + mobileNumber
}

func getTokenExiries() (time.Duration, time.Duration) {
	accessTokenExpiry, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	refreshTokenExpiry, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY"))

	return time.Duration(accessTokenExpiry), time.Duration(refreshTokenExpiry)
}

func getClaims(token string) (jwt.MapClaims, error) {
	secretKey := []byte(os.Getenv("JWT_TOKEN"))
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func GetUserIdFromToken(token string) (int, error) {
	claims, err := getClaims(token)
	if err != nil {
		return 0, err
	}

	floatUserId := claims["user_id"].(float64)
	userId := int(floatUserId)

	return userId, nil
}

func GenerateToken(userId int) (*Token, error) {
	secretKey := []byte(os.Getenv("JWT_TOKEN"))
	accessTokenExpiry, refreshTokenExpiry := getTokenExiries()

	// Generating Access Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(accessTokenExpiry * time.Second).Unix(),
		"type":    ACCESS_TOKEN_TYPE,
	})
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	// Generating Refresh TOken
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(refreshTokenExpiry * time.Second).Unix(),
		"type":    REFRESH_TOKEN_TYPE,
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

func IsValidToken(token string, tokenType string) bool {
	claims, err := getClaims(token)
	if err != nil {
		return false
	}

	// checking expiry
	expiry := claims["exp"].(float64)
	expiryTime := time.Unix(int64(expiry), 0)
	currentTime := time.Now()
	if expiryTime.Before(currentTime) {
		return false
	}

	// check if user exists
	floatUserId := claims["user_id"].(float64)
	userId := int(floatUserId)
	_, err = models.GetUser(userId)
	if err != nil {
		return false
	}

	// check token type
	tokenTypeFromToken := claims["type"].(string)

	return tokenTypeFromToken == tokenType
}

func IsRateLimitExceeded(mobileNumber string) bool {
	key := getOTPAttempPrefix(mobileNumber)
	val, err := helpers.GetCacheValue(key)
	if err != nil {
		return false
	}
	num, _ := strconv.Atoi(val)
	fmt.Println(key, num, val)

	return num > OTP_MAX_ATTEMPT
}

func IncrementOTPAttemptCounter(mobileNumber string) {
	key := getOTPAttempPrefix(mobileNumber)
	val, err := helpers.GetCacheValue(key)
	if err != nil {
		helpers.SetCacheValue(key, "1", OTP_ATTEMPT_KEY_EXPIRY)
		return
	}
	num, _ := strconv.Atoi(val)
	num++
	fmt.Println(key, num, val)
	helpers.SetCacheValue(key, strconv.Itoa(num), OTP_ATTEMPT_KEY_EXPIRY)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request headers
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		if !IsValidToken(tokenString, ACCESS_TOKEN_TYPE) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		} else {
			userId, _ := GetUserIdFromToken(tokenString)
			c.Set("user_id", userId) // Set the user ID in the context for later use
			c.Next()
		}
	}

}
