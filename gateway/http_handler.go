package main

import (
	"errors"
	"net/http"

	"github.com/friday1602/common"
	pb "github.com/friday1602/common/api"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := validateItems(items); err != nil {
		common.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})

	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.ResponseWithError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		common.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.ResponseWithJSON(w, http.StatusOK, order)
}

func validateItems(items []*pb.ItemWithQuantity) error {
	if len(items) == 0 {
		return errors.New("items must have at least one item")
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("item's ID is required")
		}

		if i.Quantity <= 0 {
			return errors.New("item has invalid quantity")
		}
	}
	return nil
}
