package service

import (
	"auth2/app/dao/auth"
	"auth2/app/model"
	"auth2/utils/logger"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) ODICProvider(data model.OpenIDConfig) error {
	err := auth.ODICProviderCreate(data)
	if err != nil {
		logger.Logger().Error(err)
		return err
	}
	return nil
}
