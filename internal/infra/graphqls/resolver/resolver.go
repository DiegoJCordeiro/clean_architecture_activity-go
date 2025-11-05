package resolver

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/usecases"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/graphqls/models"
)

// Resolver é o resolver principal
type Resolver struct {
	CreateOrderUseCase *usecases.CreateOrderUseCase
	ListOrdersUseCase  *usecases.ListOrdersUseCase
}

// NewResolver cria um novo resolver
func NewResolver(
	createUseCase *usecases.CreateOrderUseCase,
	listUseCase *usecases.ListOrdersUseCase,
) *Resolver {
	return &Resolver{
		CreateOrderUseCase: createUseCase,
		ListOrdersUseCase:  listUseCase,
	}
}

// Query retorna o resolver de queries
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation retorna o resolver de mutations
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

// QueryResolver interface
type QueryResolver interface {
	ListOrders(ctx context.Context) ([]*models.Order, error)
}

// MutationResolver interface
type MutationResolver interface {
	CreateOrder(ctx context.Context, input models.CreateOrderInput) (*models.Order, error)
}

// ListOrders lista todas as ordens
func (r *queryResolver) ListOrders(ctx context.Context) ([]*models.Order, error) {
	output, err := r.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var orders []*models.Order
	for _, order := range output {
		orders = append(orders, &models.Order{
			ID:         order.ID,
			CustomerID: order.CustomerID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
		})
	}

	return orders, nil
}

// CreateOrder cria uma nova ordem
func (r *mutationResolver) CreateOrder(ctx context.Context, input models.CreateOrderInput) (*models.Order, error) {
	useCaseInput := dtos.CreateOrderInputDTO{
		CustomerID: input.CustomerID,
		Price:      input.Price,
		Tax:        input.Tax,
	}

	output, err := r.CreateOrderUseCase.Execute(ctx, useCaseInput)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		ID:         output.ID,
		CustomerID: output.CustomerID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
		CreatedAt:  output.CreatedAt,
	}, nil
}
