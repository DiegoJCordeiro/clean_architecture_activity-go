package main

import (
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/services"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/application/usecases"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/configuration"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/graphqls"
	pb "github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/grpc/protobuff"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/repositories"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/web/handlers"
	"github.com/DiegoJCordeiro/clean-architecture-activity-go/internal/infra/web/webserver"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config := configuration.NewConfiguration()

	config, err := config.Load("app", "env", ".")

	if err != nil {
		log.Fatal("Error loading configuration. ", err)
	}

	db, err := configuration.ConnectMongoDB(config.MongoDBHost, config.MongoDBDatabase)

	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	orderRepository := repositories.NewOrderRepository(db)

	createOrderUseCase := usecases.NewCreateOrderUseCase(orderRepository)
	listOrdersUseCase := usecases.NewListOrdersUseCase(orderRepository)

	go startRESTServer(config.WebServerPort, createOrderUseCase, listOrdersUseCase)
	go startGRPCServer(config.GrpcPort, createOrderUseCase, listOrdersUseCase)

	startGraphQLServer(config.GraphQLPort, createOrderUseCase, listOrdersUseCase)
}

func startRESTServer(port string, createUC *usecases.CreateOrderUseCase, listUC *usecases.ListOrdersUseCase) {

	webServer := webserver.NewWebServer(port)
	webServer.AddMiddleware()

	orderHandler := handlers.NewOrderHandler(createUC, listUC)

	webServer.Router.Post("/order", orderHandler.CreateOrder)
	webServer.Router.Get("/order", orderHandler.ListOrders)

	if err := webServer.Start(); err != nil {
		log.Fatalf("Erro ao iniciar REST API: %v", err)
	}
}

func startGRPCServer(port string, createUC *usecases.CreateOrderUseCase, listUC *usecases.ListOrdersUseCase) {

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Erro ao iniciar listener gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderService := services.NewOrderService(createUC, listUC)

	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	log.Printf("🚀 gRPC Server rodando na porta %s\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro ao iniciar gRPC server: %v", err)
	}
}

func startGraphQLServer(port string, createUC *usecases.CreateOrderUseCase, listUC *usecases.ListOrdersUseCase) {
	resolverInstance := graphqls.NewResolver(createUC, listUC)
	graphqlHandler := graphqls.CreateGraphQLServer(resolverInstance)

	http.Handle("/graphql", graphqlHandler)

	log.Printf("🚀 GraphQL Server rodando na porta %s\n", port)
	log.Printf("   GraphiQL disponível em: http://localhost:%s/graphql\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Erro ao iniciar GraphQL server: %v", err)
	}
}
