package api

import (
	"github.com/gin-gonic/gin"
	"kfonline/service/user"
	"kfonline/util/request"
)

func LoadUserRoute(r *gin.Engine) {
	g := r.Group("/user")
	//检查是否有可用的在线客服
	g.GET("check/online/kf", func(c *gin.Context) {
		user.CheckOnlineKf(request.New(c))
	})
	//发送消息
	g.POST("publish", func(c *gin.Context) {
		user.Publish(request.New(c))
	})
	//获取连接凭证
	g.GET("jwt", func(c *gin.Context) {
		user.Jwt(request.New(c))
	})

}
