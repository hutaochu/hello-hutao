package main

import (
	"github.com/hutaochu/hello-hutao/web"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func main() {
	web.RunApp()
}
