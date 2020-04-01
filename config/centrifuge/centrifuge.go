package centrifuge

import (
	"github.com/centrifugal/gocent"
	"kfonline/config/env"
)

var (
	isLoad   = false
	instance = new(gocent.Client)
)

func connection() {
	instance = gocent.New(gocent.Config{
		Addr:       env.Centrifuge.Addr,
		Key:        env.Centrifuge.ApiKey,
		HTTPClient: nil,
	})
	isLoad = true
}

func Instance() *gocent.Client {
	if !isLoad {
		connection()
	}
	return instance
}
