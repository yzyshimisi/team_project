package userService

import (
	"TRS/app/models"
	"TRS/config/database"

	"gorm.io/gorm"
)

func CheckPhoneNumExist(phoneNum string) error {
	result := database.DB.Where("phone_num=?", phoneNum).First(&models.User{})
	return result.Error
}

func GetUserByUserName(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username=?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CheckPass(pass1 string, pass2 string) bool {
	return pass1 == pass2
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}

func ComparePhoneNum(phoneNum1 string, phoneNum2 string) bool {
	return phoneNum1 == phoneNum2
}

func CheckUserNameExist(username string) bool {
	result := database.DB.Where("username=?", username).First(&models.User{})
	if result.Error == gorm.ErrRecordNotFound {
		return true
	}
	return false
}
