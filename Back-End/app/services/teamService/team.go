package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CheckOwnerNameExist(username string) error {
	result := database.DB.Where("username=?", username).First(&models.Team{})
	return result.Error
}

func CheckTeamExistByTeamID(team_id string) error {
	result := database.DB.Where("team_id=?", team_id).First(&models.Team{})
	return result.Error
}

func CreateTeam(team models.Team) error {
	result := database.DB.Create(&team)
	return result.Error
}

func GetTeamInfoByTeamID(team_id string) (*models.Team, error) {
	var team models.Team
	result := database.DB.Where("team_id=?", team_id).First(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

func CheckPass(pass1 string, pass2 string) bool {
	return pass1 == pass2
}

func JoinTeam(team models.Team) error {
	result := database.DB.Create(&team)
	return result.Error
}

func UpdateTeamNum(team_id string, num int) error {
	result := database.DB.Model(&models.Team{}).Where("team_id=?", team_id).Update("team_num", num)
	return result.Error
}

func CheckRight(username string) error {
	result := database.DB.Where("captain_name=?", username).First(&models.Team{})
	return result.Error
}

func DeleteTeamMember(kicked_name string) error {
	result := database.DB.Delete(models.Team{}, "username=?", kicked_name)
	return result.Error
}

func DeleteTeam(captain_name string) error {
	result := database.DB.Delete(&models.Team{}, "captain_name=?", captain_name)
	return result.Error
}

func UpdateTeamStatus(captain_name string, new_status string) error {
	result := database.DB.Model(&models.Team{}).Where("captain_name=?", captain_name).Update("status", new_status)
	return result.Error
}

func GetTeamMember(team_id string) (*[]string, error) {
	var members *[]string
	result := database.DB.Model(&models.Team{}).Where("team_id=?", team_id).Pluck("username", &members) //Pluck 用于从数据库查询单个列，并将结果扫描到切片
	return members, result.Error
}

func GetTeamInfoByCaptainName(captain_name string) (*models.Team, error) {
	var team models.Team
	result := database.DB.Where("captain_name=?", captain_name).First(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}
