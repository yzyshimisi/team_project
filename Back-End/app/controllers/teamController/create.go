package teamController

import (
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"
	"TRS/config/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTeamData struct {
	TeamName string `json:"team_name" binding:"required"`
	TeamID   string `json:"team_id" binding:"required"`
	TeamPass string `json:"team_pass" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

func CreateTeam(c *gin.Context) {
	//创建团队
	var data CreateTeamData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断创建者是否已加入团队
	err = teamService.CheckOwnerNameExist(data.UserName)
	if err == nil {
		utils.JsonErrorResponse(c, 200507, "用户已有团队")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断团队编号是否存在
	err = teamService.CheckTeamExistByTeamID(data.TeamID)
	if err == nil {
		utils.JsonErrorResponse(c, 200508, "团队编号已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//创建团队
	err = teamService.CreateTeam(models.Team{
		TeamID:      data.TeamID,
		TeamName:    data.TeamName,
		TeamPass:    data.TeamPass,
		TeamNum:     1,
		CaptainName: data.UserName,
		Status:      "uncommitted",
		UserName:    data.UserName,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	var team *models.Team
	result := database.DB.Where("team_id=?", data.TeamID).First(&team)
	if result.Error != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, team)
}
