package router

import (
	"auth2/utils/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	register(r)
	return r
}
