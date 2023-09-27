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
	authService *service.AuthService
}

func NewAuth() Auth {
	return Auth{
		authService: service.NewAuthService(),
	}
}

// ODICProvider odic 服务商
// @Summary 服务商注册接口
// @Description 服务商注册接口
// @Tags 服务商相关接口
// @Accept application/json
// @Produce application/json
// @Param data body service.OpenIDConfigRequest true  "服务商注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} httpResp.Resp
// @Router /auth/odic_providers/register [post]
func (a *Auth) ODICProviderRegister(c *gin.Context) *httpResp.Response {
	var data model.OpenIDConfig
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return httpResp.HttpResp(code.ParamsError, err.Error())
		}
		return httpResp.HttpResp(code.ParamsError, errs.Translate(trans))
	}
	err := a.authService.ODICProviderRegister(data)
	if err != nil {
		return httpResp.HttpResp(code.DBError, err.Error())
	}
	return httpResp.HttpResp(code.Success)
}

// ODICProvider odic 服务商
// @Summary 获取生成的应用凭证
// @Description 获取生成的应用凭证
// @Tags 服务商相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} httpResp.Resp
// @Router /auth/applications/credentials [get]
func (a *Auth) ProviderApplicationGenerateCredentials(c *gin.Context) *httpResp.Response {
	data := a.authService.ProviderApplicationGenerateCredentials()
	return httpResp.HttpResp(code.Success, data)

}

// ODICProvider odic 服务商
// @Summary 服务商应用注册
// @Description 服务商应用注册
// @Tags 服务商相关接口
// @Accept application/json
// @Produce application/json
// @Param data body service.ProviderApplicationRequest true  "服务商应用注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} httpResp.Resp
// @Router /auth/applications/register [post]
func (a *Auth) ProviderApplicationRegister(c *gin.Context) *httpResp.Response {
	var data model.ProviderApplication
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return httpResp.HttpResp(code.ParamsError, err.Error())
		}
		return httpResp.HttpResp(code.ParamsError, errs.Translate(trans))
	}
	err := a.authService.ProviderApplicationRegister(data)
	if err != nil {
		return httpResp.HttpResp(code.DBError, err.Error())
	}
	return httpResp.HttpResp(code.Success)
}

// ODICProvider odic 服务商
// @Summary 获取服务商应用
// @Description 获取服务商应用
// @Tags 服务商相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} httpResp.Resp
// @Router /auth/applications [get]
func (a *Auth) ProviderApplicationList(c *gin.Context) *httpResp.Response {
	data := a.authService.ProviderApplicationList()
	return httpResp.HttpResp(code.Success, data)
}
