package controller

import (
	"auth2/app/service"
	"auth2/utils/code"
	"auth2/utils/httpResp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OIDC struct {
	oidcService *service.OIDCService
}

func NewOIDC() OIDC {
	return OIDC{
		oidcService: service.NewOIDCService(),
	}
}

func (oidc *OIDC) Authorize(c *gin.Context) *httpResp.Response {
	var data service.AuthorizeRequest
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return httpResp.HttpResp(code.ParamsError, err.Error())
		}
		return httpResp.HttpResp(code.ParamsError, errs.Translate(trans))
	}
	url := oidc.oidcService.Authorize(data)
	c.Redirect(http.StatusMovedPermanently, url)
	return httpResp.HttpResp()
}

func (ocid *OIDC) GrantAndRedirect(c *gin.Context) *httpResp.Response {
	var data service.GrantAndRedirectRequest
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return httpResp.HttpResp(code.ParamsError, err.Error())
		}
		return httpResp.HttpResp(code.ParamsError, errs.Translate(trans))
	}
	ocid.oidcService.GrantAndRedirect(data)
	return httpResp.HttpResp()
}
