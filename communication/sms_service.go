package communication

import (
	"fmt"
	"strings"
)

type SMSHelper interface {
	SendSMS() error
}

type OTPBodyStruct struct {
	To  string
	OTP string
}

func (b *OTPBodyStruct) SendSMS() error {
	smsBody, err := GetSMSBody("OTP")
	if err != nil {
		return err
	}

	finalSMS := strings.Replace(smsBody, "<OTP>", b.OTP, 1)
	fmt.Println(finalSMS, b.To)

	return nil
}
