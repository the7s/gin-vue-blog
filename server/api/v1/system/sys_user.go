package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/the7s/go-vue-blog/server/model/system"
	systemReq "github.com/the7s/go-vue-blog/server/model/system/request"
)

type UserApi struct{}

func (u UserApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.BindJSON(&r)

	user := &system.User{Username: r.Username, Password: r.Password}
	_, _ = userService.Register(*user)

	c.JSON(0, "这是用户注册接口")
}

func (u UserApi) Login(c *gin.Context) {

	userName, _ := c.Params.Get("userName")
	fmt.Println(userName)
}
