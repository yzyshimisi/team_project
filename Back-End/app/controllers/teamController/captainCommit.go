package teamController

import (
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamCommitData struct {
	UserName  string `json:"username" binding:"required"`
	NewStatus string `json:"new_status" binding:"required"`
}

func TeamCommit(c *gin.Context) {
	var data TeamCommitData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断是否为队长
	err = teamService.CheckRight(data.UserName)
	if err != nil && err == gorm.ErrRecordNotFound {
		utils.JsonErrorResponse(c, 200510, "无授权")
		return
	} else if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//获取团队信息
	var team *models.Team
	team, err = teamService.GetTeamInfoByCaptainName(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//更新状态
	if team.Status == data.NewStatus {
		utils.JsonErrorResponse(c, 200511, "已提交报名")
		return
	}

	if team.TeamNum < 4 {
		utils.JsonErrorResponse(c, 200512, "团队人数不够")
		return
	}

	err = teamService.UpdateTeamStatus(data.UserName, data.NewStatus)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, data.NewStatus)
}
