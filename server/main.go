package main

import (
	"github.com/the7s/go-vue-blog/server/core"
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://gpproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod dowload

func main() {
	global.GVB_VP = core.Viper() // 初始化配置
	global.GVB_LOG = core.Zap()
	global.GVB_DB = initialize.Gorm()

	core.RunServer()
}
