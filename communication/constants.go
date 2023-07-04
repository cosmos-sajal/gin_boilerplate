package communication

import "errors"

func GetSMSBody(key string) (string, error) {
	SMSBodyMap := make(map[string]string)
	SMSBodyMap["OTP"] = "Your OTP is <OTP>"

	value, ok := SMSBodyMap[key]
	if !ok {
		return "", errors.New("Invalid SMS body key")
	} else {
		return value, nil
	}
}
