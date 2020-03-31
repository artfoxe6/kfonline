package api

import (
	"github.com/gin-gonic/gin"
	"kfonline/middleware"
)

func LoadKfRoute(r *gin.Engine) {

	kf := r.Group("/kf")
	kf.Use(middleware.Auth())
	{
		// 获取uuid ,通过id唯一生成，确保每次获取的一致
		kf.GET("/uuid", func(c *gin.Context) {

		})

		//发送消息
		kf.POST("/publish", func(c *gin.Context) {

		})

		//客服登录
		kf.POST("/login", func(c *gin.Context) {

		})

		//获取以往的聊天列表
		kf.GET("/history", func(c *gin.Context) {

		})
	}

}
