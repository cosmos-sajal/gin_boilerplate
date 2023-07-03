package otpservice

import (
	"fmt"
	"strconv"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
)

func getOTPAttempPrefix(mobileNumber string) string {
	return OTP_ATTEMPTS_PREFIX + mobileNumber
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
