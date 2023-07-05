package otpservice

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos-sajal/go_boilerplate/communication"
	"github.com/cosmos-sajal/go_boilerplate/helpers"
	queueservice "github.com/cosmos-sajal/go_boilerplate/queue_service"
)

func getOTPAttempPrefix(mobileNumber string) string {
	return OTP_ATTEMPTS_PREFIX + mobileNumber
}

func SendOTP(mobileNumber string, otp string, userId int) {
	shouldUseQueueingSystem, err :=
		strconv.ParseBool(os.Getenv("SHOULD_USE_QUEUEING_SYSTEM"))
	if err != nil || !shouldUseQueueingSystem {
		otpBodyStruct := &communication.OTPBodyStruct{
			OTP: otp,
			To:  mobileNumber,
		}
		go otpBodyStruct.SendSMS()

		return
	}

	smsSenderStruct := queueservice.SendSMSStruct{
		OTP:          otp,
		MobileNumber: mobileNumber,
		UserId:       userId,
	}
	smsSenderStruct.SendMessage()
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
