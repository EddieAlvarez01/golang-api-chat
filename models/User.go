package models

import "time"

type User struct {
	ID        uint64
	RoleID    uint64
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//TableName es para el nombre de la tabla en mysql
func (User) TableName() string {
	return "User"
}
