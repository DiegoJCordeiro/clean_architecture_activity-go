package handlers

import (
	"encoding/json"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/usecases"
	ports "github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/adapters/usecases"
	"net/http"
)

type OrderHandler struct {
	CreateOrderUseCase ports.CreateOrdersUseCasePort
	ListOrdersUseCase  ports.ListOrdersUseCasePort
}

// NewOrderHandler cria um novo handler
func NewOrderHandler(
	createUseCase *usecases.CreateOrderUseCase,
	listUseCase *usecases.ListOrdersUseCase,
) *OrderHandler {
	return &OrderHandler{
		CreateOrderUseCase: createUseCase,
		ListOrdersUseCase:  listUseCase,
	}
}

// CreateOrder cria uma nova ordem via REST
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateOrderInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// ListOrders lista todas as ordens via REST
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListOrdersUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
