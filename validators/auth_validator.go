package validators

import (
	"net/http"

	"github.com/cosmos-sajal/go_boilerplate/models"
	authservice "github.com/cosmos-sajal/go_boilerplate/services/auth"
)

type SignInStruct struct {
	MobileNumber *string `json:"mobile_number"`
}

func (s *SignInStruct) Validate() *RestError {
	if s.MobileNumber == nil {
		return &RestError{
			Status: http.StatusNotFound,
			Errors: []Error{{
				Key:     "mobile_number",
				Message: "Mobile Number is required",
			}},
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

	return nil
}

type RefreshTokenStruct struct {
	RefreshToken *string `json:"refresh_token"`
}

func (r *RefreshTokenStruct) Validate() *RestError {
	if r.RefreshToken == nil {
		return &RestError{
			Status: http.StatusBadRequest,
			Errors: []Error{{
				Key:     "refresh_token",
				Message: "Refresh Token is required",
			}},
		}
	}

	if !authservice.IsValidToken(*r.RefreshToken, authservice.REFRESH_TOKEN_TYPE) {
		return &RestError{
			Status: http.StatusBadRequest,
			Errors: []Error{{
				Key:     "refresh_token",
				Message: "Invalid Refresh Token",
			}},
		}
	}

	return nil
}
