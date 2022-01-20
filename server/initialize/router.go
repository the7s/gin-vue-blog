package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/the7s/go-vue-blog/server/router"
)

// 初始化总路由

func Routers(r *gin.Engine) {

	systemGroup := router.RouterGroupApp.System
	// 健康监测
	PublicGroup := r.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemGroup.InitUserRouter(PublicGroup) //注册用户路由
	}

}
