package api

import (
	"github.com/gin-gonic/gin"
	"kfonline/middleware"
	"kfonline/service/kf"
	"kfonline/util/request"
)

func LoadKfRoute(r *gin.Engine) {

	g := r.Group("/kf")
	g.Use(middleware.Auth())
	{
		//添加客服
		g.POST("/add", func(c *gin.Context) {
			kf.AddKf(request.New(c))
		})
		// 删除客服
		g.POST("/del", func(c *gin.Context) {
			kf.DelKf(request.New(c))
		})
		//发送消息
		g.POST("/publish", func(c *gin.Context) {
			kf.Publish(request.New(c))
		})
		//客服登录
		g.POST("/login", func(c *gin.Context) {
			kf.Login(request.New(c))
		})
		//获取连接凭证
		g.POST("/jwt", func(c *gin.Context) {
			kf.Jwt(request.New(c))
		})
		//获取私有频道连接凭证
		g.POST("/jwt/private", func(c *gin.Context) {
			kf.PrivateJwt(request.New(c))
		})
	}

}
