package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"

	"github.com/tal-tech/go-zero/zrpc"
)

//增加服务
type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder     //增加添加的服务
	Checker checker.Checker //增加检查的服务

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
