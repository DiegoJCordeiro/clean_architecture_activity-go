package services

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/usecases"
	ports "github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/adapters/usecases"
	protobuff "github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/grpc/protobuff"
)

type OrderService struct {
	protobuff.UnimplementedOrderServiceServer
	CreateOrderUseCase ports.CreateOrdersUseCasePort
	ListOrdersUseCase  ports.ListOrdersUseCasePort
}

func NewOrderService(
	createUseCase *usecases.CreateOrderUseCase,
	listUseCase *usecases.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createUseCase,
		ListOrdersUseCase:  listUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *protobuff.CreateOrderRequest) (*protobuff.CreateOrderResponse, error) {
	input := dtos.CreateOrderInputDTO{
		CustomerID: req.CustomerId,
		Price:      req.Price,
		Tax:        req.Tax,
	}

	output, err := s.CreateOrderUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &protobuff.CreateOrderResponse{
		Id:         output.ID,
		CustomerId: output.CustomerID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
		CreatedAt:  output.CreatedAt,
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *protobuff.ListOrdersRequest) (*protobuff.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var orders []*protobuff.Order
	for _, order := range output {
		orders = append(orders, &protobuff.Order{
			Id:         order.ID,
			CustomerId: order.CustomerID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
		})
	}

	return &protobuff.ListOrdersResponse{
		Orders: orders,
	}, nil
}
