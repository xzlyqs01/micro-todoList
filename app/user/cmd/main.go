package main

import (
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"micro-todoList/micro-todoList/app/user/repository/db/dao"
	"micro-todoList/micro-todoList/app/user/service"
	"micro-todoList/micro-todoList/config"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"micro-todoList/micro-todoList/idl/pb"
)

func main() {
	config.Init()
	dao.InitDB()
	//etcd注册
	etcdReg:=etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s",config.EtcdHost,config.EtcdPort)))
	//得到微服务实例
	microService:=micro.NewService(
		micro.Name("rpcUserAddress"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_=pb.RegisterUserServiceHandler(microService.Server(),service.GetUserSrv())

	//启动微服务
	_=microService.Run()

}
