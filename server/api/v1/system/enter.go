package system

import "github.com/the7s/go-vue-blog/server/service"

type ApiGroup struct {
	JwtApi
	UserApi
}

var (
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
