package rdb

import (
	"context"
	"auth2/config"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	RDB *redis.Client
)

func InitRedis()error{
	RDB = redis.NewClient(&redis.Options{
        Addr:     config.RedisConfig.Addr,
        Password: config.RedisConfig.Password,
        DB:       int(config.RedisConfig.DB),
		PoolSize: int(config.RedisConfig.PoolSize),
		MinIdleConns: int(config.RedisConfig.MinIdleConns),
    })
	ctx := context.Background()
	_,err := RDB.Ping(ctx).Result()
	if err != nil{
		return err
	}
	logrus.Println("config mysql inited... ...")
	return nil
}