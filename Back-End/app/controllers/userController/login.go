package userController

import (
	"TRS/app/models"
	"TRS/app/services/userService"
	"TRS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` //binding的作用就是如果前端没有给我相应的数据，打回
}

func Login(c *gin.Context) {
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//获取用户信息
	var user *models.User
	user, err = userService.GetUserByUserName(data.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	//判断密码是否正确
	flag := userService.CheckPass(data.Password, user.Password)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "密码错误")
		return
	}

	utils.JsonSuccessResponse(c, user)
}
