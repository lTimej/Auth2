package router

import (
	"auth2/config"

	"github.com/gin-gonic/gin"
)

func Router(middlewares ...gin.HandlerFunc) *gin.Engine {
	var r *gin.Engine
	if config.ServerConfig.Mode == "dev" {
		r = gin.Default()
	} else {
		r = gin.New()
	}
	if len(middlewares) > 0 {
		r.Use(middlewares...)
	}
	return r
}
