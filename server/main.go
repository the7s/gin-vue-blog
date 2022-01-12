package main

import (
	"github.com/the7s/go-vue-blog/server/core"
	"github.com/the7s/go-vue-blog/server/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://gpproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod dowload

func main() {
	global.GVB_VP = core.Viper() // 初始化配置
	core.RunServer()
}
