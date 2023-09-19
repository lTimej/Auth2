package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	ServerConfig *ServerConf
	MysqlConfig  *MysqlConf
	RedisConfig  *RedisConf
	LoggerConfig *LoggerConf
	ZapConfig    *ZapConf
)

type ServerConf struct {
	Host string
	Port int
	Name string
	Mode string
}

type MysqlConf struct {
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
}

type RedisConf struct {
	Addr         string
	PoolSize     uint32
	MinIdleConns uint32
	Password     string
	DB           uint32
}

type LoggerConf struct {
	Level       string
	FilePath    string
	FileName    string
	MaxFileSize uint64
	ToFile      bool
}

type ZapConf struct {
	Level      string
	Filename   string
	MaxSize    uint64
	MaxBackups uint64
	MaxAge     uint64
}

func initServerConf() {
	ServerConfig = &ServerConf{
		Host: viper.GetString("server.host"),
		Port: viper.GetInt("server.port"),
		Name: viper.GetString("server.name"),
		Mode: viper.GetString("server.mode"),
	}
}

func initMysqlConf() {
	MysqlConfig = &MysqlConf{
		DataSourceName: viper.GetString("mysql.dataSourceName"),
		MaxOpenConns:   viper.GetInt("mysql.maxOpenConns"),
		MaxIdleConns:   viper.GetInt("mysql.maxIdleConns"),
	}
	fmt.Println(MysqlConfig, "99999")
}

func initRedisConf() {
	RedisConfig = &RedisConf{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetUint32("redis.db"),
		PoolSize:     viper.GetUint32("redis.poolSize"),
		MinIdleConns: viper.GetUint32("redis.minIdleConns"),
	}
}

func initLoggerConf() {
	LoggerConfig = &LoggerConf{
		Level:       viper.GetString("logger.level"),
		FilePath:    viper.GetString("logger.filePath"),
		FileName:    viper.GetString("logger.fileName"),
		MaxFileSize: viper.GetUint64("logger.maxFileSize"),
		ToFile:      viper.GetBool("logger.toFile"),
	}
}

func initZapConf() {
	ZapConfig = &ZapConf{
		Level:      viper.GetString("zap.level"),
		Filename:   viper.GetString("zap.fileName"),
		MaxSize:    viper.GetUint64("zap.maxSize"),
		MaxAge:     viper.GetUint64("zap.maxAge"),
		MaxBackups: viper.GetUint64("zap.maxBackups"),
	}
}

func InitConf() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	fmt.Println(3333)
	initServerConf()
	initMysqlConf()
	initRedisConf()
	initLoggerConf()
	initZapConf()
	return nil
}
