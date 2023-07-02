package models

import (
	"time"

	"github.com/cosmos-sajal/go_boilerplate/helpers"
	"github.com/cosmos-sajal/go_boilerplate/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	DOB          time.Time
	MobileNumber string
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index"`
}

type UserResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	UUID         uuid.UUID `json:"uuid"`
	DOB          time.Time `json:"dob"`
	MobileNumber string    `json:"mobile_number"`
}

func CreateUser(name string, mobileNumber string, dob string) (*UserResponse, error) {
	datetimeDob, err := helpers.ConvertStringToDatetime(dob)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:         name,
		MobileNumber: mobileNumber,
		DOB:          datetimeDob,
	}

	result := initializers.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		MobileNumber: user.MobileNumber,
		UUID:         user.UUID,
		DOB:          user.DOB,
	}, nil
}

type UserList struct {
	Users []UserResponse `json:"users"`
	Count int64          `json:"count"`
}

func GetUserList(limit int, offset int) (*UserList, error) {
	var users []User
	var resultCount int64
	var userResponses []UserResponse

	result := initializers.DB.Model(&users).Where(
		"deleted_at is null").Limit(limit).Offset(offset).Find(&userResponses)
	if result.Error != nil {
		return nil, result.Error
	}

	initializers.DB.Model(&User{}).Where("deleted_at is null").Count(&resultCount)

	return &UserList{
		Users: userResponses,
		Count: resultCount,
	}, nil
}

func IsNumberPresent(MobileNumber string) bool {
	result := initializers.DB.Where(
		"mobile_number = ? and deleted_at is null", MobileNumber).First(&User{})

	return result.RowsAffected > 0
}
