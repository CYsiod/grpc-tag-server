package main

import (
	"flag"
	pb "github.com/CYsiod/grpc-tag-server/proto"
	"github.com/CYsiod/grpc-tag-server/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
