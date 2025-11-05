package usecases

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/adapters/repositories"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/models"
)

type CreateOrderUseCase struct {
	OrderRepository repositories.OrderRepositoryPort
}

func NewCreateOrderUseCase(repository repositories.OrderRepositoryPort) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: repository,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input dtos.CreateOrderInputDTO) (*dtos.CreateOrderOutputDTO, error) {

	order, err := models.NewOrder(input.CustomerID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = uc.OrderRepository.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateOrderOutputDTO{
		ID:         order.ID.Hex(),
		CustomerID: order.CustomerID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
		CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
