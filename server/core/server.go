package core

import (
	"fmt"
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	address := fmt.Sprintf(":%d", global.GVB_CONFIG.System.Addr)
	fmt.Printf(`
	欢迎使用 github.com/the7s/gin-vue-blog/server
	当前版本:V0.0.1 beta
	默认自动化文档地址:http://127.0.0.1%s`, address)

	Router := initialize.Routers()

	s := initServer(address, Router)

	s.ListenAndServe().Error()
}
