package auth

import (
	"auth2/app/dao/db"
	"auth2/app/model"
)

func ODICProviderCreate(op model.OpenIDConfig) error {
	if err := db.DB.Create(&op).Error; err != nil {
		return err
	}
	return nil
}

func GetODICProviderByProvideCode(providerCode string) *model.OpenIDConfig {
	var oidc *model.OpenIDConfig
	db.DB.Where("sched_provider_code = ?", providerCode).First(oidc)
	return oidc
}

func GetProviderApplicationByCode(ProviderCode string) *model.ProviderApplication {
	var pa model.ProviderApplication
	db.DB.Where("provider_code = ?", ProviderCode).First(&pa)
	return &pa
}

func GetProviderApplicationByClienId(client_id string) *model.ProviderApplication {
	var pa model.ProviderApplication
	db.DB.Where("client_id = ?", client_id).First(&pa)
	return &pa
}

func ProviderApplicationCreate(pa model.ProviderApplication) error {
	if err := db.DB.Create(&pa).Error; err != nil {
		return err
	}
	return nil
}

func GetProviderApplicationAll() []model.ProviderApplication {
	var pa []model.ProviderApplication
	db.DB.Find(&pa)
	return pa
}
