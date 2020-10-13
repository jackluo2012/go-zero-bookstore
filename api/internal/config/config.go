package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

//通过etcd自动去发现可用的add/check服务
type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf //添加 etcd 的add 发现服务
	Check zrpc.RpcClientConf //添加 check 的 发现服务
}
