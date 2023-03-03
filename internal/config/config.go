package config

import (
	"log"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	LogLevel     int
}

type EnvConfig struct {
	MysqlConfig
}

var envConfig *EnvConfig

func InitConfig() {
	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigName(GetEnv())
	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("init config error: ", err)
	}

	envConfig = &EnvConfig{
		MysqlConfig: MysqlConfig{
			Host:         v.GetString("mysql.host"),
			Port:         v.GetString("mysql.port"),
			User:         v.GetString("mysql.user"),
			Password:     v.GetString("mysql.password"),
			DatabaseName: v.GetString("mysql.database_name"),
			LogLevel:     v.GetInt("mysql.log_level"),
		},
	}
}

func GetEnvConfig() *EnvConfig {
	return envConfig
}
