package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDb() {
	db, err := gorm.Open(sqlite.Open("nayu.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	DB = db
}

func GetDB() *gorm.DB {
	if DB == nil {
		initDb()
	}

	return DB
}
