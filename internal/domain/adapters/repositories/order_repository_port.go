package repositories

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/models"
)

type OrderRepositoryPort interface {
	Create(ctx context.Context, order *models.Order) error
	FindAll(ctx context.Context) ([]*models.Order, error)
	FindByID(ctx context.Context, id string) (*models.Order, error)
}
