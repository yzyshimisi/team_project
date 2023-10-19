package teamController

import (
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JoinTeamData struct {
	TeamID   string `json:"team_id" binding:"required"`
	TeamPass string `json:"team_pass" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

func JoinTeam(c *gin.Context) {
	var data JoinTeamData
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200509, "团队不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	//获取团队信息
	var team *models.Team
	team, err = teamService.GetTeamInfoByTeamID(data.TeamID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断团队是否满员
	if team.TeamNum == 6 {
		utils.JsonErrorResponse(c, 200513, "团队已满员")
		return
	}
	//判断团队是否处于提交状态
	if team.Status == "committed" {
		utils.JsonErrorResponse(c, 200514, "团队已报名")
		return
	}
	//判断团队密码是否正确
	flag := teamService.CheckPass(data.TeamPass, team.TeamPass)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "密码错误")
		return
	}
	//加入团队
	err = teamService.JoinTeam(models.Team{
		TeamID:      team.TeamID,
		TeamName:    team.TeamName,
		TeamNum:     team.TeamNum,
		TeamPass:    team.TeamPass,
		CaptainName: team.CaptainName,
		Status:      team.Status,
		UserName:    data.UserName,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//更新团队人数
	num := team.TeamNum + 1
	err = teamService.UpdateTeamNum(team.TeamID, num)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	team, err = teamService.GetTeamInfoByTeamID(data.TeamID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, team)
}
