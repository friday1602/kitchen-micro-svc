package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/friday1602/common/api"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("HTTP_PORT")
	orderServiceAddr := os.Getenv("ORDER_SERVICE_ADDR")

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

	log.Println("listening to port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)

	}
}
