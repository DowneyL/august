package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	configName = "config"
	configPath = "config"
	configType = "yml"
	conf       *viper.Viper
)

func Init(args ...string) {
	if len(args) > 0 {
		configPath = filepath.Clean(args[0])
	}

	conf = viper.New()

	conf.AddConfigPath(configPath)
	conf.SetConfigType(configType)
	conf.SetConfigName(configName)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Conf() *viper.Viper {
	return conf
}

func SetConf(key string, value interface{}) {
	conf.Set(key, value)
}

func Filepath() string {
	return fmt.Sprintf("%s/%s.%s", configPath, configName, configType)
}
