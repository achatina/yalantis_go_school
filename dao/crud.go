package dao

import "gorm.io/gorm"

func LogSession() {

}

func GetSessionNumbers(db *gorm.DB) int64 {
	var count int64
	db.Table("session").Count(&count)
	return count
}