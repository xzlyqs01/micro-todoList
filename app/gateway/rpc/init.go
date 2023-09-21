package rpc

import (
	"go-micro.dev/v4"
	"micro-todoList/micro-todoList/idl/pb"
)

var(
	UserService pb.UserService

)

func InitRPC(){
	userMicroService:=micro.NewService(
		micro.Name("userService.client"),
		//micro.WrapClient(wrappers.NewUserWrapper),
		)
	//用户服务调度实例
	userService:=pb.NewUserService("rpcUserService",userMicroService.Client())


	UserService=userService
}





