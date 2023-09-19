package db

import (
	"os"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"auth2/config"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
	err error
)

func InitMysql()error{
	l := log.New(os.Stdout, "\r\n", log.LstdFlags)
	newLogger := logger.New(l,logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})
	DB, err = gorm.Open(mysql.Open(config.MysqlConfig.DataSourceName), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Println("init mysql failed... ...", err)
		return err
	}
	logrus.Println("config mysql inited... ...")
	return nil
}