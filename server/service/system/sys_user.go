package system

import (
	//"errors"
	"fmt"
	//"github.com/jinzhu/gorm"
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/model/system"
)

type UserService struct{}

func (UserService *UserService) Register(u system.User) (err error, userInter system.User) {
	var user system.User

	err = global.GVB_DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	u.Password = "124321312"
	err = global.GVB_DB.Create(&u).Error
	return err, u
}
