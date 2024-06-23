package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/friday1602/common"
	pb "github.com/friday1602/common/api"
)

var (
	httpAddr         = common.EnvString("HTTP_PORT", ":8080")
	orderServiceAddr = "localhost:2000" 
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Dialing orders service at", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	r := chi.NewRouter()

	handler := NewHandler(c)
	handler.registerRoutes(r)

	log.Println("listening to port", httpAddr)
	if err := http.ListenAndServe(httpAddr, r); err != nil {
		log.Fatal(err)

	}
}
