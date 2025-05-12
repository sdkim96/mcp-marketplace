package db

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/sdkim96/mcp-marketplace/internal/models"
)

type UserTable struct {
	ID       string `gorm:"primaryKey"`
	UserName string `gorm:"not null"`
	Email    string `gorm:"unique,not null"`
	Password string `gorm:"not null"`
}

// Get user entity by ID
//
// This function's SQL CMD: **SELECT * FROM user WHERE id = ? LIMIT 1;**
func GetUserByID(db *gorm.DB, id string) (*UserTable, error) {
	user := &UserTable{}
	err := db.
		Where("id = ?", id).
		First(user).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get user entity by Email
//
// This function's SQL CMD: **SELECT * FROM user WHERE email = ? LIMIT 1;**
func GetUserByEmail(db *gorm.DB, email string) (*UserTable, error) {
	user := &UserTable{}
	err := db.
		Where("email = ?", email).
		First(user).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Add new user entity
//
// This function's SQL CMD: **INSERT INTO user (id, username, email) VALUES (?, ?, ?);**
func AddUser(db *gorm.DB, singupReqModel *models.SignupRequest) error {

	user := &UserTable{}
	user.ID = uuid.New().String()
	user.UserName = singupReqModel.UserName
	user.Email = singupReqModel.Email
	user.Password = singupReqModel.Password

	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
