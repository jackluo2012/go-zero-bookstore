package main

import (
	"context"
	"log"

	"bookstore/test/pb"

	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
)

func main() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "hello.rpc",
		},
	})

	conn := client.Conn()
	hello := pb.NewGreeterClient(conn)
	reply, err := hello.SayHello(context.Background(), &pb.HelloRequest{Name: "go-zero"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.Message)
}
