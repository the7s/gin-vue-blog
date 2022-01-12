package global

import (
	"github.com/spf13/viper"
	"github.com/the7s/go-vue-blog/server/config"
)

var (
	GVB_CONFIG config.Server
	GVB_VP     *viper.Viper
)
