package http

import (
	"github.com/gin-gonic/gin"
	"micro-todoList/micro-todoList/idl/pb"
	"micro-todoList/micro-todoList/pkg/ctl"
	"net/http"
)

func UserRegisterHAndler(ctx gin.Context){
	var req pb.UserRequest
	if err:=ctx.ShouldBind(&req);err!=nil{
		ctx.JSON(http.StatusBadRequest,ctl.RespError(ctx,err,""))
	}
}











