package userController

import (
	"TRS/app/models"
	"TRS/app/services/userService"
	"TRS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	Username        string `json:"username" binding:"required"`
	Sex             string `json:"sex" binding:"required"`
	StudentNum      string `json:"student_num" binding:"required"`
	PhoneNum        string `json:"phone_num" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断手机号是否已被注册
	err = userService.CheckPhoneNumExist(data.PhoneNum)
	if err == nil {
		utils.JsonErrorResponse(c, 200504, "手机号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断用户名是否已存在
	flag := userService.CheckUserNameExist(data.Username)
	if !flag {
		utils.JsonErrorResponse(c, 200513, "用户名已存在")
		return
	}
	//确认两次输入的密码是否一致
	flag = userService.CheckPass(data.Password, data.ConfirmPassword)
	if !flag {
		utils.JsonErrorResponse(c, 200505, "密码不一致")
		return
	}

	err = userService.Register(models.User{
		Username:   data.Username,
		Password:   data.Password,
		Sex:        data.Sex,
		StudentNum: data.StudentNum,
		PhoneNum:   data.PhoneNum,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	user, err := userService.GetUserByUserName(data.Username)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, user)
}
