package user

import (
	"context"
	"github.com/satori/go.uuid"
	"kfonline/config/centrifuge"
	"kfonline/config/env"
	"kfonline/model/MessageModel"
	"kfonline/util/request"
	"kfonline/util/token"
	"time"
)

//获取共有jwt
func Jwt(r *request.Request) bool {
	userUuid := uuid.NewV4()
	jwt, _ := token.CreateJwtToken(map[string]interface{}{
		"sub": userUuid,
		"exp": time.Now().Add(time.Second * time.Duration(env.Jwt.Exp)).Unix(),
	})
	return r.Success(map[string]interface{}{
		"uuid": userUuid,
		"jwt":  jwt,
	})
}

//往频道发送消息
func Publish(r *request.Request) bool {
	clams, err := r.TokenClams()
	if err != nil {
		return r.Error(err.Error())
	}
	userUuid := clams["uuid"].(string)
	message := &MessageModel.Message{
		User:    userUuid,
		Message: r.Post("message"),
		At:      time.Now(),
	}
	err = message.Create()
	if err != nil {
		return r.Error(err.Error())
	}
	_ = centrifuge.Instance().Publish(context.Background(), r.Post("channel"), []byte(r.Post("message")))
	return r.Success(nil)
}

//用户订阅自己的私有频道时获取私有jwt验证
func PrivateJwt(r *request.Request) bool {
	var form struct {
		Client   string   `form:"client" binding:"required"`
		Channels []string `form:"channels" binding:"required"`
	}
	if err := r.C.ShouldBind(&form); err != nil {
		return r.Error(err.Error())
	}
	clams, err := r.TokenClams()
	if err != nil {
		return r.Error(err.Error())
	}
	userUuid := clams["uuid"].(string)
	if ("$" + userUuid) != form.Channels[0] {
		return r.Error("subscribe failed")
	}
	res, err := token.CreateJwtToken(map[string]interface{}{
		"client":  form.Client,
		"channel": form.Channels[0],
	})
	if err != nil {
		return r.Error(err.Error())
	}
	resp := map[string][]map[string]string{
		"channels": {
			{
				"channel": form.Channels[0],
				"token":   res,
			},
		},
	}
	return r.Success(resp)
}

func CheckOnlineKf(r *request.Request) bool {
	res, err := centrifuge.Instance().PresenceStats(context.Background(), env.Kf.Channel)
	if err != nil {
		return r.Error(err.Error())
	}
	if res.NumUsers == 0 {
		return r.Error("当前没有客服在线")
	}

	return r.Success(nil)
}
