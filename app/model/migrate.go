package model

import (
	"auth2/app/dao/db"
	"auth2/utils/logger"
)

func MakeMigrations() error {
	if err := db.DB.AutoMigrate(&User{}, &Client{}, &Token{}, &Provider{}, &ProviderApplication{}, &OpenIDConfig{}); err != nil {
		logger.Logger().Error("数据迁移失败：", err)
		return err
	}
	logger.Logger().Info("数据迁移成功")
	return nil
}
