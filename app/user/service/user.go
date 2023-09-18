package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"micro-todoList/micro-todoList/app/user/repository/db/dao"
	"micro-todoList/micro-todoList/app/user/repository/db/model"
	"micro-todoList/micro-todoList/idl/pb"
	"micro-todoList/micro-todoList/pkg/e"
	"sync"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv()*UserSrv{
	UserSrvOnce.Do(func() {
		UserSrvIns=&UserSrv{}
	})
	return UserSrvIns
}

func (u*UserSrv) UserLogin(ctx context.Context,req*pb.UserRequest,resp*pb.UserDetailResponse)(err error){
	resp.Code= e.SUCCESS
	user,err:=dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err!=nil{
		resp.Code=e.ERROR
		return
	}
	if !user.CheckPassword(req.Password){
		resp.Code=e.InvalidParams
		return
	}
	resp.UserDetail=BuildUser(user)
	return
}

func (u*UserSrv) UserRegister(ctx context.Context,req*pb.UserRequest,resp*pb.UserDetailResponse)(err error){
	if req.Password!=req.PasswordConfirm{
		err=errors.New("两次密码输入不一致")
		return
	}
	resp.Code=e.SUCCESS
	_,err=dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err!=nil{
		if err==gorm.ErrRecordNotFound{

		}else{
			resp.Code=e.ERROR
			return
		}
	}
	user:=&model.User{
		UserName: req.UserName,
	}
	if err=user.SetPassword(req.Password);err!=nil{
		resp.Code=e.ERROR
		return
	}
	if err=dao.NewUserDao(ctx).CreateUser(user);err!=nil{
		resp.Code=e.ERROR
		return
	}
	resp.UserDetail=BuildUser(user)
	return
}


func BuildUser(item *model.User)*pb.UserModel{
	UserModel:=pb.UserModel{
		Id: uint32(item.ID),
		UserName: item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &UserModel
}





