package router

import (
	"github.com/gin-gonic/gin"
	"auth2/utils/httpResp"
)



type handlerFunc func(*gin.Context) *httpResp.Response

func response(h handlerFunc)gin.HandlerFunc {
	return func(c *gin.Context){
		r := h(c)
		if r != nil{
			c.JSON(r.HttpStatus, &r.Result)
		}
		print(111111)
		httpResp.PutResponse(r)
	}
}