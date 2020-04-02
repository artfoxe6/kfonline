package MarkModel

import (
	"kfonline/config/db"
	"kfonline/model/KfModel"
	"time"
)

type Mark struct {
	ID      uint       `gorm:"primary_key" json:"id"`
	User    string     `json:"user" gorm:"size:40"`
	KfUid   uint       `json:"kf_uid"`
	Content string     `json:"content"`
	At      time.Time  `json:"at"`
	Kefu    KfModel.Kf `gorm:"foreignkey:KfUid" json:"kefu"`
}

func (Mark) TableName() string {
	return "mark"
}

func (m *Mark) Create() error {
	return db.Instance().Create(m).Error
}

type MarkList []Mark

func (list *MarkList) List(uuid string, page, perPage int) error {
	return db.Instance().Preload("Kefu").
		Where("user=?", uuid).Offset((page - 1) * perPage).Limit(perPage).Find(list).Error
}
