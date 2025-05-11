package db

type UserTable struct {
	ID       string `gorm:"primaryKey"`
	UserName string
	Email    string
}
