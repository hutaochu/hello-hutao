package config

import (
	"flag"

	"github.com/hutaochu/hello-hutao/constants"
)

var env string

func init() {
	flag.StringVar(&env, "env", constants.DevEnviroment, "application env")
	flag.Parse()
}

func GetEnv() string {
	return env
}
