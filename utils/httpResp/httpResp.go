package httpResp

import (
	"reflect"
	"net/http"
	// "fmt"
	code2 "auth2/utils/code"
	"sync"
)


type Resp struct{
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code uint32 `json:"code"`
}

type Response struct{
	HttpStatus int
	Result Resp
}

var pool = sync.Pool{
	New: func()interface{}{
		return &Response{}
	},
}

func NewResponseOk(status int, code uint32, data interface{})*Response{
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.Result.Code = code
	response.Result.Data = data
	response.Result.Msg = code2.GetMeg(code)
	return response
}

func NewResponseNotOk(status int, code uint32, msg string)*Response{
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.Result.Code = code
	response.Result.Data = nil
	response.Result.Msg = msg
	return response
}

func PutResponse(res *Response){
	if res != nil{
		res.Result.Data = nil
		pool.Put(res)
	}
}

func HttpResp(params ...interface{})*Response{
	var code uint32
	var msg string
	var data interface{}
	for _,param := range(params){
		t := reflect.TypeOf(param).Kind()
		switch t {
		case reflect.Uint32:
			code = param.(uint32)
		case reflect.String:
			msg = param.(string)
		default:
			data = param
		}
	}
	if msg == ""{
		return NewResponseOk(http.StatusOK, code, data)
	}else{
		return NewResponseNotOk(http.StatusOK, code, msg)
	}
}