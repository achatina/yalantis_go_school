package dao

type Session struct {
	Id    int    `json:"id" gorm:"id"`
	Ip    string `json:"ip" gorm:"ip"`
	Token string `json:"token" gorm:"token"`
}
