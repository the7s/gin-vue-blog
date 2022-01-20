package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/the7s/go-vue-blog/server/api/v1"
)

type UserRouter struct{}

func (ur UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("register", userApi.Register) // 注册用户
	}

	return userRouter
}
