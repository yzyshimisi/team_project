package userController

import (
	"TRS/app/models"
	"TRS/app/services/userService"
	"TRS/app/utils"
	"TRS/config/database"

	"github.com/gin-gonic/gin"
)

type ChangePassData struct {
	UserName        string `json:"username"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ChangePass(c *gin.Context) {
	var data ChangePassData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200500, "参数错误")
		return
	}
	//获取用户信息
	var user *models.User
	user, err = userService.GetUserByUserName(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断两次密码是否一致
	flag := userService.CheckPass(data.NewPassword, data.ConfirmPassword)
	if !flag {
		utils.JsonErrorResponse(c, 200505, "密码不一致")
		return
	}
	//更新密码
	result := database.DB.Model(&user).Update("password", data.NewPassword)
	if result.Error != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
