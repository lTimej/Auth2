package validators

import (
	"auth2/app/dao/auth"

	"github.com/go-playground/validator/v10"
)

func ProviderCodeValidation(fl validator.FieldLevel) bool {
	provider_code := fl.Field().String()
	provider_code_obj := auth.GetProviderApplicationByCode(provider_code)
	if provider_code_obj == nil {
		return false
	}
	return true
}

func InitValidator() error {
	validate := validator.New()

	// 注册方法
	err := validate.RegisterValidation("providerCodeValidation", ProviderCodeValidation)
	if err != nil {
		return err
	}
	return nil
}
