package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {

}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(r *chi.Mux) {
	r.HandleFunc("/", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {

}