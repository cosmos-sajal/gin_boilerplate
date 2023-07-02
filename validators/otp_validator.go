package validators

import (
	"net/http"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/models"
	authservice "github.com/cosmos-sajal/go_boilerplate/services"
)

type OTPValidatorStruct struct {
	MobileNumber *string `json:"mobile_number"`
	OTP          *string `json:"otp"`
}

func (s *OTPValidatorStruct) Validate() *RestError {
	var errors []Error

	authservice.IncrementOTPAttemptCounter(*s.MobileNumber)
	if s.MobileNumber == nil {
		errors = append(errors, Error{
			Key:     "mobile_number",
			Message: "Mobile Number is required",
		})
	}

	if s.OTP == nil {
		errors = append(errors, Error{
			Key:     "otp",
			Message: "OTP is required",
		})
	}

	if len(errors) > 0 {
		return &RestError{
			Status: http.StatusNotFound,
			Errors: errors,
		}
	}

	if !models.IsNumberPresent(*s.MobileNumber) {
		return &RestError{
			Status: http.StatusBadRequest,
			Errors: []Error{{
				Key:     "mobile_number",
				Message: "Mobile Number is not registered",
			}},
		}
	}

	otp, err := helpers.GetCacheValue("OTP_KEY_" + *s.MobileNumber)
	if err != nil {
		return &RestError{
			Status: http.StatusBadRequest,
			Errors: []Error{{
				Key:     "otp",
				Message: "OTP is expired",
			}},
		}
	}

	if authservice.IsRateLimitExceeded(*s.MobileNumber) {
		return &RestError{
			Status: http.StatusTooManyRequests,
			Errors: []Error{{
				Key:     "otp",
				Message: "OTP limit exceeded",
			}},
		}
	}

	if otp != *s.OTP {
		return &RestError{
			Status: http.StatusBadRequest,
			Errors: []Error{{
				Key:     "otp",
				Message: "Invalid OTP",
			}},
		}
	}

	return nil
}
