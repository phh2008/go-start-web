package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
)

var ConfigSet = wire.NewSet(NewConfig)

// ConfigFolder 配置文件目录
type ConfigFolder string

type Config struct {
	ConfigDir ConfigFolder
	Viper     *viper.Viper
}

func NewConfig(configFolder ConfigFolder) *Config {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(string(configFolder))
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("加载配置错误")
		panic(err)
	}
	return &Config{
		ConfigDir: configFolder,
		Viper:     viper,
	}
}

func (a *Config) GetString(key string) string {
	return a.Viper.GetString(key)
}

func (a *Config) Get(key string) interface{} {
	return a.Viper.Get(key)
}
