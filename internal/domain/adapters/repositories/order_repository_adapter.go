package repositories

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/entities"
)

type OrderRepositoryPort interface {
	Create(ctx context.Context, order *entities.Order) error
	FindAll(ctx context.Context) ([]*entities.Order, error)
	FindByID(ctx context.Context, id string) (*entities.Order, error)
}
