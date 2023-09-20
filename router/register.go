package router

import (
	"auth2/app/controller"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	auth_controller := controller.NewAuth()
	auth_group := router.Group("/auth")
	{
		auth_group.GET("/authorize", response(auth_controller.Authorize))
		auth_group.POST("/odic_providers/register", response(auth_controller.ODICProvider))
	}
}
