package validators

import (
	"fmt"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/models"
)

type CreateUser struct {
	Name         *string `json:"name"`
	MobileNumber *string `json:"mobile_number"`
	Dob          *string `json:"dob"`
}

func (u *CreateUser) Validate() *RestError {
	var errors []Error

	if u.Name == nil {
		errors = append(errors, Error{
			Key:     "name",
			Message: "Required Field",
		})
	}

	if u.MobileNumber == nil {
		errors = append(errors, Error{
			Key:     "mobile_number",
			Message: "Required Field",
		})
	}

	if len(errors) > 0 {
		return &RestError{
			Status: 400,
			Errors: errors,
		}
	}

	if !helpers.ValidateMobileNumber(*u.MobileNumber) {
		errors = append(errors, Error{
			Key:     "mobile_number",
			Message: "Invalid Mobile Number",
		})
	}

	if models.IsNumberPresent(*u.MobileNumber) {
		errors = append(errors, Error{
			Key:     "mobile_number",
			Message: "Already present",
		})
	}

	if u.Dob != nil && !helpers.ValidateDateOfBirth(*u.Dob) {
		errors = append(errors, Error{
			Key:     "dob",
			Message: "Invalid Date Of Birth",
		})
	}

	if len(errors) > 0 {
		return &RestError{
			Status: 400,
			Errors: errors,
		}
	}

	return nil
}

type GetUserList struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

func (u *GetUserList) Validate() *RestError {
	var errors []Error

	if u.Limit != nil && *u.Limit < 0 {
		errors = append(errors, Error{
			Key:     "limit",
			Message: "Invalid Limit",
		})
	} else if u.Limit == nil {
		u.Limit = new(int)
		*u.Limit = 10
		fmt.Println("idhar aaya")
	}

	if u.Offset != nil && *u.Offset < 0 {
		errors = append(errors, Error{
			Key:     "offset",
			Message: "Invalid Offset",
		})
	} else if u.Offset == nil {
		u.Offset = new(int)
		*u.Offset = 0
		fmt.Println("idhar aaya 2")
	}

	if len(errors) > 0 {
		return &RestError{
			Status: 400,
			Errors: errors,
		}
	}

	return nil
}
