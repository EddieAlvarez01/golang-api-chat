package models

import "time"

type Message struct {
	ID        uint      `json:"id"`
	Receiver  User      `json:"receiver"`
	Sender    User      `json:"sender"`
	Message   string    `json:"message"`
	Active    uint8     `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

//TableName PARA INDICAR EL NOMBRE DE LA TABLA
func (Message) TableName() string {
	return "Message"
}
