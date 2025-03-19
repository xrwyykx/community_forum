package global

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		Port     string
		DBName   string
		Charset  string
	}
	HTTP struct {
		Path string
		Port string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")   //确定配置文件的名称
	viper.SetConfigType("toml")     //确定配置文件的类型
	viper.AddConfigPath("./config") //在指定的目录中查找配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failerd to read config", err)
	}
	//将配置内容解析到全局变量AppConfig中
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("failed to unmarshal config")
	}
	log.Println("config loaded successfully")
}
