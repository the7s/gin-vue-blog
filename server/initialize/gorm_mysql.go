package initialize

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/the7s/go-vue-blog/server/global"
)

func GormMysql() *gorm.DB {
	m := global.GVB_CONFIG.Mysql

	if m.Dbname == "" {
		return nil
	}

	mysql.NewConfig()
	db, err := gorm.Open("mysql", m.Dsn())
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return global.GVB_CONFIG.Mysql.TablePrefix + defaultTableName
	}
	return db
}
