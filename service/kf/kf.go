package kf

import (
	"context"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"kfonline/config/centrifuge"
	"kfonline/config/env"
	"kfonline/model/KfModel"
	"kfonline/model/MessageModel"
	"kfonline/util/lib"
	"kfonline/util/request"
	"kfonline/util/token"
	"time"
)

//获取共有jwt
func Jwt(r *request.Request) bool {
	kfUuid := uuid.NewV3(uuid.UUID{}, r.Post("phone")).String()
	jwt, _ := token.CreateJwtToken(map[string]interface{}{
		"sub": kfUuid,
		"exp": time.Now().Add(time.Second * time.Duration(env.Jwt.Exp)).Unix(),
		"kf":  1,
	})
	return r.Success(map[string]interface{}{
		"uuid": kfUuid,
		"jwt":  jwt,
	})
}

//登录
func Login(r *request.Request) bool {
	if err := r.Validate([]string{"phone", "password"}); err != nil {
		return r.Error(err.Error())
	}
	kf := new(KfModel.Kf)
	err := kf.Find(r)
	if err != nil {
		return r.Error("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(kf.Password), []byte(r.Post("password")))
	if err != nil {
		return r.Error("密码错误")
	}
	kfUuid := uuid.NewV3(uuid.UUID{}, r.Post("phone")).String()
	return r.Success(map[string]interface{}{
		"uuid": kfUuid,
		//"history": history(kfuuid, r),
	})
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
	kfUuid := []byte(r.Post("channel"))[1:]
	message := &MessageModel.Message{
		User:    string(kfUuid),
		Message: r.Post("message"),
		At:      time.Now(),
	}
	err := message.Create()
	if err != nil {
		return r.Error(err.Error())
	}
	_ = centrifuge.Instance().Publish(context.Background(), r.Post("channel"), []byte(r.Post("message")))
	return r.Success(nil)
}

//获取历史消息记录
func history(kfuuid string, r *request.Request) {
	//kfs := new(KfModel.Kfs)
	//err := kfs.List(kfuuid, r.Page(), r.PerPage())
	//if err != nil {
	//	return nil
	//}
	//return kfs
}

//订阅用户的私有频道时获取私有jwt验证
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
	isKf := clams["kf"].(string)
	if isKf != "1" {
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
