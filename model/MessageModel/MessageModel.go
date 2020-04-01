package MessageModel

import (
	"kfonline/config/db"
	"kfonline/model/KfModel"
	"time"
)

type Message struct {
	ID      uint       `gorm:"primary_key" json:"id"`
	User    string     `json:"user"`
	Kf      string     `json:"kf"`
	Message string     `json:"message"`
	At      time.Time  `json:"at"`
	Kefu    KfModel.Kf `gorm:"foreignkey:Kf" json:"kefu"`
}

func (m *Message) Create() error {
	return db.Instance().Create(m).Error
}

type MessageList []Message

func (list *MessageList) ListForKf(uuid string, page, perPage int) error {
	return db.Instance().Preload("Kefu").
		Where("user=?", uuid).Offset((page - 1) * perPage).Limit(page).Find(list).Error
}

func (list *MessageList) ListForUser(uuid string, page, perPage int) error {
	return db.Instance().Where("user=?", uuid).Offset((page - 1) * perPage).Limit(page).Find(list).Error
}
