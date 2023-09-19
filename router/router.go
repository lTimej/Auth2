package router

import (
	"github.com/gin-gonic/gin"
)

func Router()*gin.Engine {
	r := gin.Default()

	register(r)
	return r
}
