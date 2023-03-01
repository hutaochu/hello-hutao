package mysql

import (
	"fmt"
	"log"

	"github.com/hutaochu/hello-hutao/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	mysqlConf := config.GetEnvConfig().MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DatabaseName,
	)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(mysqlConf.LogLevel)),
	})
	if err != nil {
		log.Fatal("create mysql connect error: ", err)
	}
	db = conn
}

func GetMysqlClient() *gorm.DB {
	return db
}
