package main

import (
	"auth2/app/dao/db"
	"auth2/app/dao/rdb"
	"auth2/config"
	"auth2/router"
	"auth2/utils/logger"
	"auth2/utils/middlewares"
	"fmt"
	"strconv"
)

func main() {
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

	engine := router.Router(middlewares.GinRecovery(true))
	router.Register(engine)
	addr := config.ServerConfig.Host + ":" + strconv.Itoa(config.ServerConfig.Port)
	engine.Run(addr)
}
