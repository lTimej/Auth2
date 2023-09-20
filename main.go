package main

import (
	"auth2/app/dao/db"
	"auth2/app/dao/rdb"
	"auth2/app/model"
	"auth2/config"
	"auth2/router"
	"auth2/utils/logger"
	"auth2/utils/middlewares"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var db_flag string
	flag.StringVar(&db_flag, "m", "", "数据迁移")
	flag.Parse()

	config.InitConf()

	if err := logger.InitLogger(); err != nil {
		panic(fmt.Errorf("load logger failed, reason:%s", err.Error()))
	}

	if err := logger.InitZap(); err != nil {
		panic(fmt.Errorf("load zap failed, reason:%s", err.Error()))
	}

	if err := db.InitMysql(); err != nil {
		panic(fmt.Errorf("load mysql failed, reason:%s", err.Error()))
	}
	if err := rdb.InitRedis(); err != nil {
		panic(fmt.Errorf("load redis failed, reason:%s", err.Error()))
	}

	if db_flag == "migrate" {
		if err := model.MakeMigrations(); err != nil {
			panic(fmt.Errorf("db migrate failed, reason:%s", err.Error()))
		}
	}

	engine := router.Router(middlewares.GinRecovery(true))
	router.Register(engine)
	addr := config.ServerConfig.Host + ":" + strconv.Itoa(config.ServerConfig.Port)
	engine.Run(addr)
}
