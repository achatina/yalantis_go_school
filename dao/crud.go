package dao

import (
	"gorm.io/gorm"
)

func SaveSession(db *gorm.DB, session *Session) {
	db.Table("session").Create(session)
}

func UpdateSession(db *gorm.DB, session *Session) {
	db.Table("session").Save(session)
}

func GetSessionByIp(db *gorm.DB, ip string) (*Session, error) {
	var session *Session
	if err := db.Table("session").Where("ip = ?", ip).Take(&session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func GetSessionNumbers(db *gorm.DB) int64 {
	var count int64
	db.Table("session").Count(&count)
	return count
}
