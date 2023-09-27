package router

import (
	"auth2/app/controller"

	"github.com/gin-gonic/gin"

	_ "auth2/docs" // 千万不要忘了导入把你上一步生成的docs

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	// "github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	//swagger文档
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	auth_controller := controller.NewAuth()
	auth_group := router.Group("/auth")
	{
		auth_group.POST("/odic_providers/register", response(auth_controller.ODICProviderRegister))
		auth_group.GET("/applications/credentials", response(auth_controller.ProviderApplicationGenerateCredentials))
		auth_group.POST("/applications/register", response(auth_controller.ProviderApplicationRegister))
		auth_group.GET("/applications", response(auth_controller.ProviderApplicationList))
	}
	oidc_controller := controller.NewOIDC()
	oidc_group := router.Group("/oidc")
	{
		oidc_group.GET("/authorize", response(oidc_controller.Authorize))
		oidc_group.GET("/callback", response(oidc_controller.GrantAndRedirect))
	}
}
