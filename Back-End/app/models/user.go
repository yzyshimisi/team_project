package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Sex        string `json:"sex"`
	StudentNum string `json:"student_num"`
	PhoneNum   string `json:"phone_num"`
}
