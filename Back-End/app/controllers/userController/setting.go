package userController

import (
	_ "TRS/app/models"
	"TRS/app/services/userService"
	"TRS/app/utils"
	"TRS/config/database"

	"github.com/gin-gonic/gin"
)

type UpdateUserData struct {
	UserName    string `json:"username" binding:"required"`
	NewUserName string `json:"new_username"`
	Sex         string `json:"new_sex"`
	StudentNum  string `json:"new_student_num"`
	PhoneNum    string `json:"new_phone_num"`
}

func UpdateUserInformation(c *gin.Context) {
	//获取更新信息
	var data UpdateUserData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//根据操作者用户名获取原始用户信息
	user, err := userService.GetUserByUserName(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断新用户名是否存在
	flag := userService.CheckUserNameExist(data.NewUserName)
	if !flag {
		utils.JsonErrorResponse(c, 200513, "用户名已存在")
		return
	}
	//更新个人信息
	result := database.DB.Model(&user).Updates(data)
	if result.Error != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//再次获取个人信息
	user, err = userService.GetUserByUserName(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, user)

}
