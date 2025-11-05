#!/bin/bash

# Script de Instalação - Orders System
# Este script prepara o ambiente e instala todas as dependências

set -e

echo "🚀 Iniciando instalação do Orders System..."
echo ""

# Verifica se Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go não está instalado. Por favor, instale Go 1.21 ou superior."
    exit 1
fi

echo "✅ Go $(go version) detectado"

# Verifica se Docker está instalado
if ! command -v docker &> /dev/null; then
    echo "❌ Docker não está instalado. Por favor, instale Docker."
    exit 1
fi

echo "✅ Docker detectado"

# Verifica se Docker Compose está instalado
if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo "❌ Docker Compose não está instalado. Por favor, instale Docker Compose."
    exit 1
fi

echo "✅ Docker Compose detectado"

# Instala protoc se não estiver instalado
echo ""
echo "📦 Verificando protobuf compiler..."
if ! command -v protoc &> /dev/null; then
    echo "⚠️  protoc não encontrado. Instalando..."
    
    # Detecta o sistema operacional
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        sudo apt-get update
        sudo apt-get install -y protobuf-compiler
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        brew install protobuf
    else
        echo "❌ Sistema operacional não suportado para instalação automática do protoc"
        echo "   Por favor, instale manualmente: https://grpc.io/docs/protoc-installation/"
        exit 1
    fi
fi

echo "✅ protoc instalado"

# Baixa as dependências Go
echo ""
echo "📦 Baixando dependências Go..."
go mod download
go mod tidy

# Instala plugins do protoc para Go
echo ""
echo "📦 Instalando plugins protoc para Go..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adiciona $GOPATH/bin ao PATH se não estiver
export PATH="$PATH:$(go env GOPATH)/bin"

# Gera código gRPC a partir do .proto
echo ""
echo "🔨 Gerando código gRPC a partir dos arquivos .proto..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/infra/grpc/protobuff/orders.proto

echo "✅ Código gRPC gerado com sucesso"

# Cria arquivo .env se não existir
if [ ! -f .env ]; then
    echo ""
    echo "📝 Criando arquivo .env..."
    cp app.env app.env 2>/dev/null || echo "Arquivo .env já existe"
fi

echo ""
echo "✅ Instalação concluída com sucesso!"
echo ""
echo "📋 Próximos passos:"
echo "   1. Execute: docker-compose up -d (para subir o MongoDB)"
echo "   2. Execute: go run cmd/server/main.go (para iniciar a aplicação)"
echo "   3. Ou use o script de teste: ./scripts/test.sh"
echo ""
