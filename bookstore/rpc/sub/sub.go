// Code generated by goctl. DO NOT EDIT!
// Source: sub.proto

package main

import (
	"flag"
	"fmt"

	"bookstore/rpc/sub/internal/config"
	"bookstore/rpc/sub/internal/server"
	"bookstore/rpc/sub/internal/svc"
	"bookstore/rpc/sub/sub"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/sub.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewSuberServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sub.RegisterSuberServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}