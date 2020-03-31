package Migrate

import (
	"kfonline/config/db"
	"kfonline/model/KfModel"
	"kfonline/model/MarkModel"
	"kfonline/model/MessageModel"
)

func Run() {
	db.Instance().AutoMigrate(
		MarkModel.Mark{},
		MessageModel.Message{},
		KfModel.Kf{},
	)
}
