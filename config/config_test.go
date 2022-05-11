package config

import (
	"fmt"
	"testing"
)

var config *Config

func init() {
	config = NewConfig(".")
}

func Test1(t *testing.T) {

	viper := config.Viper
	fmt.Println(viper.GetString("db.url"))

}
