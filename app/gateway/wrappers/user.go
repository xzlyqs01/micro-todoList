package wrappers

import (
	"context"
	"go-micro.dev/v4/api/client"
)

type userWrapper struct {
	client.Client
}

func (wrapper *userWrapper) Call(ctx context.Context,req *client.Request,resp *interface{},opts...client.CallOption)error{

}


