package router

import (
	"auth2/utils/httpResp"
	"fmt"

	"github.com/gin-gonic/gin"
)

type handlerFunc func(*gin.Context) *httpResp.Response

func response(h handlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := h(c)
		if r != nil {
			fmt.Println(r, "=====")
			c.JSON(r.HttpStatus, &r.Result)
		}
		fmt.Println(r, "===++++==")
		httpResp.PutResponse(r)
	}
}
