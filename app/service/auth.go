package service

import (
	"auth2/app/dao/auth"
	"auth2/app/model"
	"auth2/constants"
	"auth2/utils/logger"
)

type OpenIDConfigRequest struct {
	SchedProviderCode string `json:"sched_provider_code" gorm:"type:varchar(64);unique"`
	Issuer            string `json:"issuer" gorm:"type:varchar(100)"  binding:"required"`
	ClientID          string `json:"client_id" gorm:"type:varchar(100)"  binding:"required"`
	ClientSecret      string `json:"client_secret" gorm:"type:varchar(255)"  binding:"required"`
}

type ProviderApplicationRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	Name         string `json:"name"`
	ProviderCode string `json:"provider_code"`
}
type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) ODICProviderRegister(data model.OpenIDConfig) error {
	err := auth.ODICProviderCreate(data)
	if err != nil {
		logger.Logger().Error(err)
		return err
	}
	return nil
}

func (as *AuthService) ProviderApplicationGenerateCredentials() map[string]string {
	client_id := constants.CLIENTID
	client_secret := constants.CLIENTSECRET
	return map[string]string{
		"client_id":     client_id,
		"client_secret": client_secret,
	}
}

func (as *AuthService) ProviderApplicationRegister(data model.ProviderApplication) error {
	data.UserId = 1
	data.AuthorizationGrantType = constants.GRANT_AUTHORIZATION_CODE
	data.Algorithm = constants.RS256_ALGORITHM
	data.ClientType = constants.CLIENT_CONFIDENTIAL
	err := auth.ProviderApplicationCreate(data)
	if err != nil {
		logger.Logger().Error(err)
		return err
	}
	return nil
}

func (as *AuthService) ProviderApplicationList() []model.ProviderApplication {
	pa := auth.GetProviderApplicationAll()
	return pa
}
