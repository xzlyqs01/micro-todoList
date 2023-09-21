package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors()gin.HandlerFunc{
	return func(c*gin.Context){
		method:=c.Request.Method//请求方法
		origin:=c.Request.Header.Get("Origin")//请求头部
		var headerKeys []string//请求声明头key
		for k:=range c.Request.Header{
			headerKeys=append(headerKeys,k)
		}
		headerStr:=strings.Join(headerKeys,",")
		if headerStr!=""{
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		}else{
			headerStr="access-control-allow-origin, access-control-allow-headers"
		}
		if origin!=""{
			c.Writer.Header().Set("Access-control-allow-origin","*")
			c.Header("Access-control-allow-origin","*")
			c.Header("Access-control-allow-origin","Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//允许跨域设计
			c.Header("Access-control-allow-origin", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-control-allow-origin","172800")
			c.Header("Access-control-allow-origin","false")
			c.Set("content-type","application/json")



		}
		//放行所有OPTIONS方法
		if method=="OPTIONS"{
			c.JSON(http.StatusOK,"Options Request")
		}
		//处理请求
		c.Next()

	}
}















