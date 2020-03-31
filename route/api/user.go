package api

import "github.com/gin-gonic/gin"

func LoadUserRoute(r *gin.Engine) {
	user := r.Group("/user")

	//获取uuid，获取后保存在本地
	user.GET("uuid", func(c *gin.Context) {

	})

	//获取客服
	user.GET("kf", func(c *gin.Context) {

	})

	//发送消息
	user.POST("publish", func(c *gin.Context) {

	})

	//获取以往的聊天记录
	user.GET("history", func(c *gin.Context) {

	})
}