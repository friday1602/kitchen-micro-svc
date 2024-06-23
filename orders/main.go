package main

import (
	"context"
	"log"
	"net"

	"github.com/friday1602/common"
	"google.golang.org/grpc"
)
var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)
func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen %v",err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)	

	svc.CreateOrder(context.Background())

	log.Println("GRPC server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err !=nil {
		log.Fatal(err)
	}
}