package main

import (
	"context"
	"log"

	pb "github.com/friday1602/common/api"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received!")
	o := &pb.Order{
		Id: 42,
	}
}
