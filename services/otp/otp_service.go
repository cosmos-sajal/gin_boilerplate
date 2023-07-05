package otpservice

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/initializers"
	asyncTasks "github.com/cosmos-sajal/go_boilerplate/tasks"
	"github.com/cosmos-sajal/go_boilerplate/worker"
)

func getOTPAttempPrefix(mobileNumber string) string {
	return OTP_ATTEMPTS_PREFIX + mobileNumber
}

func SendOTP(mobileNumber string, otp string, userId int) {
	otpSMSStruct := asyncTasks.OTPSMSStruct{
		OTP:          otp,
		MobileNumber: mobileNumber,
		UserId:       userId,
	}
	stringifiedMessage, _ := json.Marshal(otpSMSStruct)
	b64EncodedReq := base64.StdEncoding.EncodeToString([]byte(stringifiedMessage))
	task := tasks.Signature{
		Name: worker.SEND_OTP_QUEUE,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: b64EncodedReq,
			},
		},
		RoutingKey: worker.SEND_OTP_QUEUE,
	}
	res, err := initializers.TaskServer.SendTask(&task)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.GetState())

	emailStruct := asyncTasks.EmailStruct{
		EmailBody: "Your OTP is " + otp,
		UserId:    userId,
	}
	stringifiedMessage, _ = json.Marshal(emailStruct)
	b64EncodedReq = base64.StdEncoding.EncodeToString([]byte(stringifiedMessage))
	task = tasks.Signature{
		Name: worker.SEND_EMAIL_QUEUE,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: b64EncodedReq,
			},
		},
		RoutingKey: worker.SEND_EMAIL_QUEUE,
	}
	res, err = initializers.TaskServer.SendTask(&task)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.GetState())
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
