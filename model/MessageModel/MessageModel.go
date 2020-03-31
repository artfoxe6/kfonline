package MessageModel

import (
	"time"
)

type Message struct {
	ID   uint      `gorm:"primary_key" json:"id"`
	User string    `json:"user"`
	Kf   string    `json:"kf"`
	At   time.Time `json:"at"`
}
