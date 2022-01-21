package system

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/model/common/response"
	"github.com/the7s/go-vue-blog/server/model/system"
	systemReq "github.com/the7s/go-vue-blog/server/model/system/request"
	systemRes "github.com/the7s/go-vue-blog/server/model/system/response"
	"github.com/the7s/go-vue-blog/server/utils"
	"go.uber.org/zap"
	"time"
)

type UserApi struct{}

// Register 用户注册
func (u UserApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.BindJSON(&r)

	user := &system.User{Username: r.Username, Password: r.Password}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.GVB_LOG.Error("注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(userReturn, c)
	}
}

// Login 用户登录
func (u UserApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.BindJSON(&l)
	user := &system.User{Username: l.Username, Password: l.Password}
	err, userReturn := userService.Login(*user)
	if err != nil {
		global.GVB_LOG.Error("用户名不存在或者密码错误", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		u.createToken(c, *userReturn)
	}
}

func (u UserApi) createToken(c *gin.Context, user system.User) {
	j := &utils.JWT{SigningKey: []byte(global.GVB_CONFIG.Jwt.SigningKey)}
	claims := systemReq.CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              //签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVB_CONFIG.Jwt.ExpiresTime, //过期时间 7天
			Issuer:    global.GVB_CONFIG.Jwt.Issuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVB_LOG.Error("获取token失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OKWithDetailed(
			systemRes.LoginResponse{
				User:  user,
				Token: token,
				ExpiresAt: global.MyTime{
					Time: time.Unix(claims.StandardClaims.ExpiresAt, 0)},
			}, "登录成功", c)
	}

}
