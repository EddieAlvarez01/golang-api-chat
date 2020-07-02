package models

import "time"

type User struct {
	ID               uint      `json:"id"`
	RoleID           uint      `json:"role_id"`
	Role             Role      `json:"role"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	MessagesReceiver []Message `gorm:"foreignkey:Receiver"`
	MessagesSender   []Message `gorm:"foreignkey:Sender"`
}

//TableName es para el nombre de la tabla en mysql
func (User) TableName() string {
	return "User"
}
