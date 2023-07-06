package tasks

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/cosmos-sajal/go_boilerplate/communication"
)

type OTPSMSStruct struct {
	OTP          string
	MobileNumber string
	UserId       int
}

func DecodeToTask(msg string, task interface{}) (err error) {
	decodedstg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return
	}
	msgBytes := []byte(decodedstg)
	err = json.Unmarshal(msgBytes, task)
	if err != nil {
		return
	}
	return
}

func SendOTP(b64payload string) error {
	payload := OTPSMSStruct{}
	DecodeToTask(b64payload, &payload)

	fmt.Println("Consuming OTP message")
	otpBody := communication.OTPBodyStruct{
		OTP: payload.OTP,
		To:  payload.MobileNumber,
	}

	return otpBody.SendSMS()
}
