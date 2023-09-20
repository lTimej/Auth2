package controller

import (
	"auth2/app/model"
	"auth2/app/service"
	"auth2/utils/code"
	"auth2/utils/httpResp"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var trans ut.Translator

type Auth struct {
	userService *service.AuthService
}

func NewAuth() Auth {
	return Auth{
		userService: service.NewAuthService(),
	}
}

func (a *Auth) Authorize(c *gin.Context) *httpResp.Response {

	return httpResp.HttpResp(code.Success, "用户名不存在", map[string]string{"msg": "time"})
}

func (a *Auth) ODICProvider(c *gin.Context) *httpResp.Response {
	var data model.OpenIDConfig
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return httpResp.HttpResp(code.ParamsError, err.Error())
		}
		return httpResp.HttpResp(code.ParamsError, errs.Translate(trans))
	}
	err := a.userService.ODICProvider(data)
	if err != nil {
		return httpResp.HttpResp(code.DBError, err.Error())
	}
	return httpResp.HttpResp(code.Success)
}
