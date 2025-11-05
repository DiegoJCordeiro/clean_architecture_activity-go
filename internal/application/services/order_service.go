package services

import (
	"context"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/dtos"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/usecases"
	ports "github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/domain/adapters/usecases"
	pb "github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/grpc/protobuff"
)

// OrderService implementa o serviço gRPC
type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase ports.CreateOrdersUseCasePort
	ListOrdersUseCase  ports.ListOrdersUseCasePort
}

// NewOrderService cria um novo serviço gRPC
func NewOrderService(
	createUseCase *usecases.CreateOrderUseCase,
	listUseCase *usecases.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createUseCase,
		ListOrdersUseCase:  listUseCase,
	}
}

// CreateOrder implementa a criação via gRPC
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	input := dtos.CreateOrderInputDTO{
		CustomerID: req.CustomerId,
		Price:      req.Price,
		Tax:        req.Tax,
	}

	output, err := s.CreateOrderUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		CustomerId: output.CustomerID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
		CreatedAt:  output.CreatedAt,
	}, nil
}

// ListOrders implementa a listagem via gRPC
func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, order := range output {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			CustomerId: order.CustomerID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
		})
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
