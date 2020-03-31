package KfModel

import (
	"kfonline/config/db"
	"kfonline/util/request"
)

type Kf struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (kf *Kf) Find(r *request.Request) error {
	return db.Instance().Where("phone=?", r.Get("phone")).First(kf).Error
}
