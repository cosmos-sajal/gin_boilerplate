package validators

import (
	"net/http"

	"github.com/cosmos-sajal/go_boilerplate/models"
)

type SignInStruct struct {
	MobileNumber *string `json:"mobile_number"`
}

func (s *SignInStruct) Validate() *RestError {
	var errors []Error

	if s.MobileNumber == nil {
		errors = append(errors, Error{
			Key:     "mobile_number",
			Message: "Mobile Number is required",
		})

		return &RestError{
			Status: http.StatusNotFound,
			Errors: errors,
		}
	}

	if !models.IsNumberPresent(*s.MobileNumber) {
		errors = append(errors, Error{
			Key:     "mobile_number",
			Message: "Mobile Number is not registered",
		})

		return &RestError{
			Status: http.StatusBadRequest,
			Errors: errors,
		}
	}

	return nil
}
