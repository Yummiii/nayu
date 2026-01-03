package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DiscordID string `gorm:"uniqueIndex"`
	Username  string
	AvatarURL string
}

func GetUserByDiscordID(db *gorm.DB, discordID string) (*User, error) {
	var user User
	result := db.Where("discord_id = ?", discordID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
