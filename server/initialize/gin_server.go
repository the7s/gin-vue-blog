package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func InitGinServer() *gin.Engine {

	// 兼容win平台 console颜色代码块乱码的情况
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	Router := gin.Default()
	return Router
}
