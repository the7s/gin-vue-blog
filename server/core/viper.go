package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/utils"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		// 解析参数
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		// 配置文件路径 传递参数 > 环境变量 > 默认值
		if config == "" {
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用默认配置%s的值,文件路径为%v\n", utils.ConfigFile, config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用环境变量%s的值,文件路径为%v\n", utils.ConfigEnv, config)
			}

		} else {
			fmt.Printf("您正在使用的命令行的-c参数传递的值,文件路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	// viper 解析 配置文件
	var v *viper.Viper = viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.GVB_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVB_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
