package teamController

import (
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"
	"TRS/config/database"

	"github.com/gin-gonic/gin"
)

type TeamMemberCheckData struct {
	TeamID   string `json:"team_id" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

func TeamMemberCheck(c *gin.Context) {
	var data TeamMemberCheckData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200500, "参数错误")
		return
	}
	//判断操作者是否属于该团队
	var team models.Team
	database.DB.Where("username=?", data.UserName).First(&team)
	if team.TeamID != data.TeamID {
		utils.JsonErrorResponse(c, 200514, "你不属于该团队")
		return
	}
	//获取队员
	var members *[]string
	members, err = teamService.GetTeamMember(data.TeamID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, members)

}
