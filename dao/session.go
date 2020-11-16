package dao

type Session struct {
	Id int `json:"id" gorm:"id"`
	Metadata string `json:"metadata" gorm:"metadata"`
}