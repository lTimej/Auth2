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
