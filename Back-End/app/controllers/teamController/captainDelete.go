package teamController

import (
	"TRS/app/services/teamService"
	"TRS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteTeamData struct {
	UserName string `json:"username" binding:"required"`
}

func DeleteTeam(c *gin.Context) {
	var data DeleteTeamData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断是否有权力修改团队信息（是否是这个团队的队长）
	err = teamService.CheckRight(data.UserName)
	if err!=nil && err==gorm.ErrRecordNotFound {
		utils.JsonErrorResponse(c, 200510, "无授权")
		return
	}else if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//删除团队
	err = teamService.DeleteTeam(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
