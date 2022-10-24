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
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yml")
	vp.AddConfigPath(string(configFolder))
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("加载配置错误")
		panic(err)
	}

	//vp.WatchConfig()
	//vp.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("config file changed:", e.Name)
	//	if err = vp.Unmarshal(&profile); err != nil {
	//		fmt.Println(err)
	//	}
	//})

	return &Config{
		ConfigDir: configFolder,
		Viper:     vp,
	}
}

func (a *Config) GetString(key string) string {
	return a.Viper.GetString(key)
}

func (a *Config) Get(key string) interface{} {
	return a.Viper.Get(key)
}
