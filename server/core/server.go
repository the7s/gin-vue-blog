package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunServer() {
	address := fmt.Sprintf(":%d", 8888)

	fmt.Printf(`
	欢迎使用 github.com/the7s/gin-vue-blog/server
	当前版本:V0.0.1 beta
	默认自动化文档地址:http://127.0.0.1%s`, address)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	err := r.Run(address)
	if err != nil {
		return
	}
}
