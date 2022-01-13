package global

import (
	"github.com/spf13/viper"
	"github.com/the7s/go-vue-blog/server/config"
	"go.uber.org/zap"
)

var (
	GVB_CONFIG config.Server
	GVB_VP     *viper.Viper
	GVB_LOG    *zap.Logger
)
