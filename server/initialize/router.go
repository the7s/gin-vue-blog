package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {

	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	Router := gin.Default()
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return Router
}
