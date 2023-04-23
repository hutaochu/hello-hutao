package entities

import (
	"fmt"
	"log"
	"sync"

	"github.com/hutaochu/hello-hutao/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var once sync.Once

// DBOpen open mysql connect
func DBOpen() {
	once = sync.Once{}
	once.Do(open)
}

func open() {
	mysqlConf := config.GetEnvConfig().MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DatabaseName,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(mysqlConf.LogLevel)),
	})
	if err != nil {
		log.Fatal("create mysql connect error: ", err)
	}
}
