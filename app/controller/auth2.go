package controller

import (
	"auth2/utils/code"
	"auth2/utils/httpResp"

	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (a *Auth) Authorize(c *gin.Context) *httpResp.Response {
	return httpResp.HttpResp(code.Success, "用户名不存在", map[string]string{"msg": "time"})
}
