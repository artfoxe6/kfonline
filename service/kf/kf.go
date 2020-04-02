package kf

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"kfonline/config/centrifuge"
	"kfonline/config/env"
	"kfonline/model/KfModel"
	"kfonline/model/MessageModel"
	"kfonline/util/lib"
	"kfonline/util/request"
	"kfonline/util/token"
	"strconv"
	"time"
)

//登录
func Login(r *request.Request) bool {
	if err := r.Validate([]string{"phone", "password"}); err != nil {
		return r.Error(err.Error())
	}
	kf := new(KfModel.Kf)
	err := kf.First(r)
	if err != nil {
		return r.Error("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(kf.Password), []byte(r.Post("password")))
	if err != nil {
		return r.Error("密码错误")
	}
	authToken, _ := token.CreateJwtToken(map[string]interface{}{
		"sub": kf.ID,
		"exp": time.Now().Add(time.Second * time.Duration(env.Jwt.Exp)).Unix(),
	})
	return r.Success(authToken)
}

//添加客服
func AddKf(r *request.Request) bool {
	if err := r.Validate([]string{"phone", "password", "name"}); err != nil {
		return r.Error(err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte(r.Post("password")), bcrypt.MinCost)
	if err != nil {
		return r.Error(err.Error())
	}
	kf := &KfModel.Kf{
		Name:     r.Post("name"),
		Phone:    r.Post("phone"),
		Password: string(password),
	}
	err = kf.Create()
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(nil)
}

//删除客服
func DelKf(r *request.Request) bool {
	if err := r.Validate([]string{"id"}); err != nil {
		return r.Error(err.Error())
	}
	m := new(KfModel.Kf)
	err := m.Del(lib.Int(r.Id()))
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(nil)
}

//往频道发送消息
func Publish(r *request.Request) bool {
	if err := r.Validate([]string{"uuid", "message"}); err != nil {
		return r.Error(err.Error())
	}
	clams, err := r.TokenClams()
	if err != nil {
		return r.Error(err.Error())
	}
	message := &MessageModel.Message{
		User:    r.Post("uuid"),
		Message: r.Post("message"),
		KfUid:   uint(lib.Int(clams["id"].(string))),
		At:      time.Now(),
	}
	err = message.Create()
	if err != nil {
		return r.Error(err.Error())
	}
	_ = centrifuge.Instance().Publish(context.Background(), r.Post("channel"), []byte(r.Post("message")))
	return r.Success(nil)
}

//客服的频道名称  默认客服的id
func KfChannel(kfId int) string {
	return strconv.Itoa(kfId)
}
