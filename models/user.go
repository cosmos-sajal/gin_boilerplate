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
	MobileNumber string    `gorm:"index:idx_mobile_number"`
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index"`
}

type UserResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	UUID         uuid.UUID `json:"uuid"`
	DOB          time.Time `json:"dob"`
	MobileNumber string    `json:"mobile_number"`
}

type UserList struct {
	Users []UserResponse `json:"users"`
	Count int64          `json:"count"`
}

type UpdateUserStruct struct {
	Name         *string
	MobileNumber *string
	DOB          *string
	IsDeleted    *bool
}

func getUserResponse(user User) *UserResponse {
	return &UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		UUID:         user.UUID,
		DOB:          user.DOB,
		MobileNumber: user.MobileNumber,
	}
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

	return getUserResponse(*user), nil
}

func GetUser(userId int) (*User, error) {
	var user User
	result := initializers.DB.Model(&User{}).Where("id = ? and deleted_at is null", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByMobile(mobileNumber string) (*User, error) {
	var user User
	result := initializers.DB.Model(&User{}).Where("mobile_number = ? and deleted_at is null", mobileNumber).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
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

func UpdateUser(userId int, userStruct *UpdateUserStruct) (*UserResponse, error) {
	user, err := GetUser(userId)
	if err != nil {
		return nil, err
	}

	if userStruct.Name != nil {
		user.Name = *userStruct.Name
	}
	if userStruct.MobileNumber != nil {
		user.MobileNumber = *userStruct.MobileNumber
	}
	if userStruct.DOB != nil {
		datetimeDob, err := helpers.ConvertStringToDatetime(*userStruct.DOB)
		if err != nil {
			return nil, err
		}
		user.DOB = datetimeDob
	}
	if userStruct.IsDeleted != nil {
		user.DeletedAt = gorm.DeletedAt{
			Time:  time.Now(),
			Valid: true,
		}
	}
	initializers.DB.Save(&user)

	return getUserResponse(*user), nil
}

func IsNumberPresent(MobileNumber string) bool {
	result := initializers.DB.Where(
		"mobile_number = ? and deleted_at is null", MobileNumber).First(&User{})

	return result.RowsAffected > 0
}

func IsUserPresent(userId int) bool {
	result := initializers.DB.Where(
		"id = ? and deleted_at is null", userId).First(&User{})

	return result.RowsAffected > 0
}
