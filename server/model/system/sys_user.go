package system

import "github.com/the7s/go-vue-blog/server/global"

type User struct {
	global.COMMON_MODEL
	Username string `json:"userName" gorm:"comment:用户登录名"`
	Password string `json:"-" gorm:"comment:用户登录密码"`
}
