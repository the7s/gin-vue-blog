package initialize

import (
	"github.com/jinzhu/gorm"
	"github.com/the7s/go-vue-blog/server/global"
)

func Gorm() *gorm.DB {
	switch global.GVB_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
