package usecases

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/adapters/repositories"
)

type ListOrdersUseCase struct {
	OrderRepository repositories.OrderRepositoryPort
}

func NewListOrdersUseCase(repository repositories.OrderRepositoryPort) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: repository,
	}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]dtos.ListOrdersOutputDTO, error) {

	orders, err := uc.OrderRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []dtos.ListOrdersOutputDTO
	for _, order := range orders {
		output = append(output, dtos.ListOrdersOutputDTO{
			ID:         order.ID.Hex(),
			CustomerID: order.CustomerID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:  order.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return output, nil
}
