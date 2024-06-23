package main

import (
	"net/http"

	"github.com/friday1602/common"
	pb "github.com/friday1602/common/api"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client} 
}

func (h *handler) registerRoutes(r *chi.Mux) {
	r.Post("/api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items: items,
	}, 	)
}
