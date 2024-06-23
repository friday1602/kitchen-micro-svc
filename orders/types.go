package main

import (
	"context"
	
	pb "github.com/friday1602/common/api"
)

type OrderService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest ) error 
}

type OrderStore interface {
	Create(context.Context) error
}