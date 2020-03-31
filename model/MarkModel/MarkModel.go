package MarkModel

import (
	"time"
)

type Mark struct {
	ID   uint      `gorm:"primary_key" json:"id"`
	User string    `json:"user"`
	At   time.Time `json:"at"`
}
