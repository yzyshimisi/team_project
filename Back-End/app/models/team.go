package models

type Team struct {
	ID          int    `json:"id"`
	TeamID      string `json:"team_id"`
	TeamName    string `json:"team_name"`
	TeamPass    string `json:"-"`
	TeamNum     int    `json:"team_num"`
	CaptainName string `json:"captain_name"`
	Status      string `json:"status"`
	UserName    string `gorm:"column:username" json:"-"`
}
