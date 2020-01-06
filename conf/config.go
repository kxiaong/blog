package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	Web struct {
		AppName      string `required:"true" flagUsage:"服务名称"`
		ListenPort   string `default:"localhost:8080" flagUsage:"服务监听地址和端口"`
		WebStaticDir string `default:"./static" flagUsage:"静态文件地址"`
	}
	Redis struct {
		Addr     string
		Password string
		DB       int `default:"0" flagUsage:"Redis数据库"`
		PoolSize int `default:"20" flagUsage:"Redis连接池大小"`
	}
	Database struct {
		Driver   string `default:"postgres" flagUsage:"数据库驱动 mysql|postgres"`
		Postgres struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
		}
		Mysql struct {
			HostPort     string `flagUsage:"数据库连接地址"`
			UserPassword string `flagUsage:"数据库账号密码"`
			DB           string `flagUsage:"数据库名"`
		}
	}

	Log struct {
		Type  string `default:"raw" flagUsage:"日志格式, json|raw"`
		Level int    `default:"5" flagUsage:"日志级别: 0 CRITICAL, 1 ERROR, 2 WARNING, 3 NOTICE, 4 INFO, 5 DEBUG"`
	} `flagUsage:"服务日志配置"`
}

var C *Config

func init() {
	current, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	ConfigFilePath := path.Join(current, "conf/config.json")
	jsonfile, err := os.Open(ConfigFilePath)
	if err != nil {
		fmt.Println("read config file failed: ", err)
		os.Exit(1)
	}
	defer jsonfile.Close()

	values, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		fmt.Println("read json file failed: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal(values, &C)
	if err != nil {
		fmt.Println("Unmarshal json file failed: ", err)
		os.Exit(1)
	}
}
