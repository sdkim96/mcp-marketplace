package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDBHandler() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../test.db"))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InitDB() {

	h := GetDBHandler()
	h.AutoMigrate(&UserTable{})

	admin := &UserTable{
		ID:       "123",
		UserName: "admin",
		Email:    "admin@sdkim.com",
	}
	err := h.Save(admin).Error
	if err != nil {
		panic("failed to add user")
	}

}
