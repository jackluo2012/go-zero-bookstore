package main

import (
	context "context"
	"flag"
	"log"

	"bookstore/test/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	grpc "google.golang.org/grpc"
)

type Config struct {
	zrpc.RpcServerConf
}

var cfgFile = flag.String("f", "./server.yaml", "cfg file")

func main() {
	flag.Parse()

	var cfg Config
	conf.MustLoad(*cfgFile, &cfg)

	srv, err := zrpc.NewServer(cfg.RpcServerConf, func(s *grpc.Server) {
		pb.RegisterGreeterServer(s, &Hello{})
	})
	if err != nil {
		log.Fatal(err)
	}
	srv.Start()
}

type Hello struct{}

func (h *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}
