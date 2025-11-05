#!/bin/bash

# Script de Testes - Orders System
# Este script testa os três endpoints: REST, gRPC e GraphQL

set -e

echo "🧪 Iniciando testes do Orders System..."
echo ""

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo ""
echo "=========================================="
echo "🔹 TESTE 1: REST API - Criar Order"
echo "=========================================="

REST_CREATE_RESPONSE=$(curl -s -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "customer-test-001",
    "price": 100.50,
    "tax": 10.05
  }')

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Order criada via REST${NC}"
    echo "Response: $REST_CREATE_RESPONSE"
else
    echo -e "${RED}❌ Erro ao criar order via REST${NC}"
fi

echo ""
echo "=========================================="
echo "🔹 TESTE 2: REST API - Listar Orders"
echo "=========================================="

REST_LIST_RESPONSE=$(curl -s -X GET http://localhost:8080/order)

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Orders listadas via REST${NC}"
    echo "Response: $REST_LIST_RESPONSE"
else
    echo -e "${RED}❌ Erro ao listar orders via REST${NC}"
fi

echo ""
echo "=========================================="
echo "🔹 TESTE 3: GraphQL - Criar Order"
echo "=========================================="

GRAPHQL_CREATE_RESPONSE=$(curl -s -X POST http://localhost:8081/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createOrder(input: { customer_id: \"customer-gql-001\", price: 150.00, tax: 15.00 }) { id customer_id price tax final_price created_at } }"
  }')

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Order criada via GraphQL${NC}"
    echo "Response: $GRAPHQL_CREATE_RESPONSE"
else
    echo -e "${RED}❌ Erro ao criar order via GraphQL${NC}"
fi

echo ""
echo "=========================================="
echo "🔹 TESTE 4: GraphQL - Listar Orders"
echo "=========================================="

GRAPHQL_LIST_RESPONSE=$(curl -s -X POST http://localhost:8081/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ listOrders { id customer_id price tax final_price created_at updated_at } }"
  }')

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Orders listadas via GraphQL${NC}"
    echo "Response: $GRAPHQL_LIST_RESPONSE"
else
    echo -e "${RED}❌ Erro ao listar orders via GraphQL${NC}"
fi

echo ""
echo "=========================================="
echo "🔹 TESTE 5: gRPC - Teste Manual"
echo "=========================================="

echo -e "${YELLOW}ℹ️  Para testar gRPC, use uma ferramenta como grpcurl:${NC}"
echo ""
echo "   # Listar Orders"
echo "   grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders"
echo ""
echo "   # Criar Order"
echo "   grpcurl -plaintext -d '{\"customer_id\": \"customer-grpc-001\", \"price\": 200.00, \"tax\": 20.00}' \\"
echo "     localhost:50051 pb.OrderService/CreateOrder"
echo ""

# Verifica se grpcurl está instalado
if command -v grpcurl &> /dev/null; then
    echo "Executando teste gRPC com grpcurl..."
    
    echo ""
    echo "Criando order via gRPC..."
    grpcurl -plaintext -d '{"customer_id": "customer-grpc-001", "price": 200.00, "tax": 20.00}' \
      localhost:50051 pb.OrderService/CreateOrder
    
    echo ""
    echo "Listando orders via gRPC..."
    grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
else
    echo -e "${YELLOW}⚠️  grpcurl não está instalado. Instale com:${NC}"
    echo "   go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest"
fi

echo ""
echo "=========================================="
echo "✅ Testes concluídos!"
echo "=========================================="
echo ""
echo "📋 Endpoints disponíveis:"
echo "   REST API:     http://localhost:8080"
echo "   gRPC Server:  localhost:50051"
echo "   GraphQL:      http://localhost:8081/graphql"
echo "   GraphiQL:     http://localhost:8081/graphql (interface web)"
echo ""
