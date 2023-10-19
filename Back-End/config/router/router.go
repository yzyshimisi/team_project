package router

import (
	"TRS/app/controllers/teamController"
	"TRS/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre1 = "/user"
	api1 := r.Group(pre1)
	{
		api1.POST("/login", userController.Login)
		api1.POST("/register", userController.Register)
		api1.PUT("/setting", userController.UpdateUserInformation)
		api1.PUT("/changepass", userController.ChangePass)
	}
	const pre2 = "/team"
	api2 := r.Group(pre2)
	{
		api2.POST("/create", teamController.CreateTeam)
		api2.POST("/join", teamController.JoinTeam)
		api2.GET("/getmember", teamController.TeamMemberCheck)
	}
	const pre3 = pre2 + "/captain"
	api3 := r.Group(pre3)
	{
		api3.PUT("/update", teamController.UpdateTeamInformation)
		api3.DELETE("/kick", teamController.KickTeamMember)
		api3.DELETE("/delete", teamController.DeleteTeam)
		api3.PUT("/commit", teamController.TeamCommit)
		api3.PUT("/backout", teamController.BackOut)
	}
}
