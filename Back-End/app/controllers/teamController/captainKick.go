package teamController

import (
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KickTeamMemberData struct {
	UserName   string `json:"username" binding:"required"`
	KickedName string `json:"kicked_name" binding:"required"`
}

func KickTeamMember(c *gin.Context) {
	var data KickTeamMemberData
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
	//获取团队信息
	var team *models.Team
	team, err = teamService.GetTeamInfoByCaptainName(data.UserName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断团队状态
	if team.Status == "committed" {
		utils.JsonErrorResponse(c, 200515, "团队已报名，不可更换队员")
		return
	}
	//删除队员
	err = teamService.DeleteTeamMember(data.KickedName)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	num := team.TeamNum - 1
	err = teamService.UpdateTeamNum(team.TeamID, num)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//返回当前队员
	var members *[]string
	members, err = teamService.GetTeamMember(team.TeamID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, members)
}
