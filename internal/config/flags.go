package config

import (
	"flag"

	"github.com/hutaochu/hello-hutao/constants"
)

var env string

func InitEnv() {
	flag.StringVar(&env, "env", constants.DevEnvironment, "application env")
	flag.Parse()
}

func GetEnv() string {
	return env
}
