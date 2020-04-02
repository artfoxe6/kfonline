package user

import (
	"context"
	"encoding/json"
	"github.com/satori/go.uuid"
	"kfonline/config/centrifuge"
	"kfonline/config/env"
	"kfonline/model/MessageModel"
	"kfonline/util/lib"
	"kfonline/util/request"
	"kfonline/util/token"
	"math/rand"
	"time"
)

//获取共有jwt
func Jwt(r *request.Request) bool {
	userUuid := uuid.NewV4()
	authToken, err := token.CreateJwtToken(map[string]interface{}{
		"sub": userUuid,
		"exp": time.Now().Add(time.Second * time.Duration(env.Jwt.Exp)).Unix(),
	})
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(authToken)
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
		KfUid:   uint(lib.Int(r.Post("channel"))),
		At:      time.Now(),
	}
	err = message.Create()
	if err != nil {
		return r.Error(err.Error())
	}
	_ = centrifuge.Instance().Publish(context.Background(), r.Post("channel"), []byte(r.Post("message")))
	return r.Success(nil)
}

//检查是否有可用的客服
func CheckOnlineKf(r *request.Request) bool {
	res, err := centrifuge.Instance().Channels(context.Background())
	if err != nil {
		return r.Error(err.Error())
	}
	if len(res.Channels) == 0 {
		return r.Error("当前没有客服在线")
	}
	channel, err := selectKf(res.Channels)
	if err != nil {
		return r.Error("暂无可用客服")
	}
	return r.Success(channel)
}

func selectKf(list []string) (string, error) {
	pipe := centrifuge.Instance().Pipe()
	for _, v := range list {
		_ = pipe.AddPresenceStats(v)
	}
	reps, err := centrifuge.Instance().SendPipe(context.Background(), pipe)
	if err != nil {
		return "", err
	}
	temp := map[string]int{}
	channels := []string{}
	for k, v := range reps {
		_ = json.Unmarshal(v.Result, &temp)
		if temp["num_users"] < env.Kf.WaitNum {
			channels = append(channels, list[k])
		}
	}
	if len(channels) == 0 {
		return "", err
	}
	rand.Seed(time.Now().UnixNano())
	return channels[rand.Intn(len(channels))], nil
}
