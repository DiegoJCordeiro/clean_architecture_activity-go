package usecases

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
)

type ListOrdersUseCasePort interface {
	Execute(ctx context.Context) ([]dtos.ListOrdersOutputDTO, error)
}

type CreateOrdersUseCasePort interface {
	Execute(ctx context.Context, input dtos.CreateOrderInputDTO) (*dtos.CreateOrderOutputDTO, error)
}
