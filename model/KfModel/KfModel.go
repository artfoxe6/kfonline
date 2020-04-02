package KfModel

import (
	"kfonline/config/db"
	"kfonline/util/request"
)

type Kf struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone" gorm:"unique size:12"`
	Password string `json:"password"`
}

func (Kf) TableName() string {
	return "kf"
}

func (kf *Kf) First(r *request.Request) error {
	return db.Instance().Where("phone=?", r.Get("phone")).First(kf).Error
}

func (kf *Kf) Create() error {
	return db.Instance().Create(kf).Error
}

func (kf *Kf) Del(id int) error {
	return db.Instance().Where("id=?", id).Delete(kf).Error
}

type Kfs []Kf

func (kfs *Kfs) List(kfuuid string, page, perpage int) error {
	return db.Instance().
		Where("kf=?", kfuuid).
		Offset((page - 1) * perpage).Limit(page).
		Find(kfs).Error
}
