package teamController

import (
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"
	"TRS/config/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateTeamData struct {
	TeamID   string `json:"new_team_id"`
	UserName string `json:"username" binding:"required"`
	TeamName string `json:"new_team_name"`
	TeamPass string `json:"new_team_pass"`
}

func UpdateTeamInformation(c *gin.Context) {
	var data UpdateTeamData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断操作者是否为队长
	err = teamService.CheckRight(data.UserName)
	if err != nil && err == gorm.ErrRecordNotFound {
		utils.JsonErrorResponse(c, 200510, "无授权")
		return
	} else if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//更新信息
	result := database.DB.Model(&models.Team{}).Where("captain_name=?", data.UserName).Updates(data)
	if result.Error != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	var team *models.Team
	team, err = teamService.GetTeamInfoByCaptainName(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, team)
}
