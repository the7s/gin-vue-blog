package system

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/model/system"
	"github.com/the7s/go-vue-blog/server/utils"
)

type UserService struct{}

func (UserService *UserService) Register(u system.User) (err error, userInter system.User) {
	var user system.User

	err = global.GVB_DB.Where("username = ?", u.Username).First(&user).Error
	fmt.Println(user)
	if !errors.Is(global.GVB_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), user
	}
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVB_DB.Create(&u).Error
	return err, u
}

func (UserService *UserService) Login(u system.User) (err error, userInter *system.User) {
	if nil == global.GVB_DB {
		return fmt.Errorf("db not init"), nil
	}
	var user system.User

	u.Password = utils.MD5V([]byte(u.Password))

	err = global.GVB_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
