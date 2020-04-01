package MarkModel

import (
	"kfonline/config/db"
	"kfonline/model/KfModel"
	"time"
)

type Mark struct {
	ID   uint       `gorm:"primary_key" json:"id"`
	User string     `json:"user"`
	Kf   string     `json:"kf"`
	At   time.Time  `json:"at"`
	Kefu KfModel.Kf `gorm:"foreignkey:Kf" json:"kefu"`
}

func (m *Mark) Create() error {
	return db.Instance().Create(m).Error
}

type MarkList []Mark

func (list *MarkList) List(uuid string, page, perPage int) error {
	return db.Instance().Preload("Kefu").
		Where("user=?", uuid).Offset((page - 1) * perPage).Limit(perPage).Find(list).Error
}
