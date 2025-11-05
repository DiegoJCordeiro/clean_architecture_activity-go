# 📦 Orders System - Clean Architecture em Go

Sistema completo de gerenciamento de pedidos implementado em **Go** seguindo os princípios de **Clean Architecture**, com suporte a **REST API**, **gRPC** e **GraphQL**.

---

## 🎯 Sobre o Projeto

Este projeto foi desenvolvido como solução para o desafio de implementação de Clean Architecture em Go, atendendo aos seguintes requisitos:

- ✅ **REST API** - Endpoints para criar e listar orders
- ✅ **gRPC** - Service com CreateOrder e ListOrders
- ✅ **GraphQL** - Mutations e Queries com interface GraphiQL
- ✅ **MongoDB** - Banco de dados NoSQL com Docker
- ✅ **Clean Architecture** - Separação em camadas (Entity, UseCase, Infrastructure)
- ✅ **Docker & Docker Compose** - Containerização completa
- ✅ **Migrations** - Scripts de inicialização do banco

---

## 🏗️ Arquitetura

O projeto segue os princípios da **Clean Architecture**, organizando o código em camadas concêntricas com dependências apontando sempre para dentro:

```
┌─────────────────────────────────────────┐
│     FRAMEWORKS & DRIVERS (Externo)     │
│  Docker, MongoDB, HTTP, gRPC, GraphQL   │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│    INTERFACE ADAPTERS (Adaptadores)     │
│  Handlers, Controllers, Presenters      │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│  APPLICATION BUSINESS RULES (Use Cases) │
│     CreateOrder, ListOrders             │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│ ENTERPRISE BUSINESS RULES (Entities)    │
│     Order, Validações, Regras           │
└─────────────────────────────────────────┘
```

### Estrutura de Diretórios

```
orders-system/
├── cmd/
│   └── clean_architecture_activity/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── domain/                  # 🔵 CAMADA 1: Domain Layer
|   │   └──adapters # Onde ficam os contratos dos Usecases e Repositories
|   │   └──entities # Onde ficam as Entidade Order com regras de negócio
│   │
│   ├── application/                 # 🟢 CAMADA 2: Application Layer
│   │   ├── usecases      # Use case de criação
│   │   |   ├── create_order_usecase.go      # Use case de criação
│   │   |   ├── list_order_usecase.go       # Use case de listagem
│   │
│   ├── infra/                   # 🟡 CAMADA 3: Infrastructure Layer
│   │   ├── database/
│   │   │   ├── mongodb.go       # Conexão MongoDB
│   │   │   └── order_repository.go
│   │   │
│   │   ├── web/                 # REST API
│   │   │   ├── handlers/
│   │   │   │   └── order_handler.go
│   │   │   └── webserver/
│   │   │       └── webserver.go
│   │   │
│   │   └── grpc/                # gRPC
│   │   │   ├── protobuff/
│   │   │   │   └── order.proto  # Definição Protocol Buffers
│   │   │   └── service/
│   │   │       └── order_service.go
│   │   │   
│   │   ├── graph/                   # GraphQL
│   │   │   └──  models/
│   │   │           └── models.go
│   │   ├── resolver/
│   │   │   ├── resolver.go
│   │   │   └── server.go
│   │   └── schema.graphql       # Schema GraphQL
│   
├── api/
│   └── api.http                 # Requisições HTTP para teste
│
├── scripts/
│   ├── test.sh                  # Script de testes
│
├── sql/
│   └── migrations/
│       └── 001_init.js          # Migração MongoDB
│
├── install.sh                   # Script de instalação
├── app.env                      # Variáveis de ambiente
├── docker-compose.yaml          # Orquestração Docker
├── Dockerfile                   # Imagem da aplicação
└── go.mod                       # Dependências Go
```

---

## 🚀 Como Executar

### Pré-requisitos

- **Go 1.24+** - [Instalar Go](https://go.dev/doc/install)
- **Docker & Docker Compose** - [Instalar Docker](https://www.docker.com/get-started)
- **Protocol Buffers (protoc)** - Para gRPC

### Opção 1: Executar Localmente (Recomendado para Desenvolvimento)

```bash
# 1. Instalar dependências
./scripts/install.sh

# 2. Subir com Docker
docker-compose up -d

# 3. Aguardar MongoDB inicializar
sleep 15

# 4. Executar aplicação
go run cmd/clean_architecture_activity/main.go
```

### Opção 2: Docker Completo

```bash
# Subir tudo com Docker
docker-compose up --build -d
```

### Opção 3: Usando Makefile

```bash
# Ver comandos disponíveis
make help

# Setup e executar
make install
make docker-up
make run
```

---

## 🌐 Endpoints e Portas

| Serviço | Porta | URL | Descrição |
|---------|-------|-----|-----------|
| **REST API** | 8080 | http://localhost:8080 | API RESTful |
| **gRPC** | 50051 | localhost:50051 | Serviço gRPC |
| **GraphQL** | 8081 | http://localhost:8081/graphql | API GraphQL + GraphiQL |
| **MongoDB** | 27017 | localhost:27017 | Banco de dados |

---

## 📡 Usando as APIs



---

### Arquivo test.sh

Use o arquivo `tests/test.sh` para testar algumas reqs graphql, grpc e rest.

### Arquivo api_test.http

Use o arquivo `tests/api_test.http` para testar algumas reqs.

---

## 🎯 Clean Architecture - Camadas

### 1. Models (Domain Layer)

**Regras de negócio puras**
**Contratos de adapters de output e usecases**

### 2. UseCase (Application Layer)

**Casos de uso da aplicação:**

### 3. Infrastructure Layer

**Implementações concretas:**

- **Repository**: Acesso ao MongoDB
- **REST Handler**: Endpoints HTTP
- **gRPC Service**: Serviço gRPC
- **GraphQL Resolver**: Queries e Mutations

---

## 🔄 Fluxo de Dados

```
Cliente (REST/gRPC/GraphQL)
        │
        ▼
    Use Case (Application)
        │
        ├─────────────┐
        ▼             ▼
    Entity      Repository
    (Domain)    (Infrastructure)
                     │
                     ▼
                 MongoDB
```

---

## 📊 Tecnologias

- **Go 1.24+** - Linguagem
- **MongoDB 7.0** - Banco de dados
- **Chi Router** - HTTP router
- **gRPC** - Protocol Buffers
- **GraphQL** - API flexível
- **Docker** - Containerização

---

## 🐛 Troubleshooting

### MongoDB não conecta

```bash
docker-compose down -v
docker-compose up -d
sleep 15
```

### Erro de autenticação

```bash
# Usar versão sem autenticação
cp app.env app.env
docker-compose up -d
```

---
## 🎓 Conceitos Aplicados

### SOLID Principles
- Single Responsibility
- Open/Closed
- Liskov Substitution
- Interface Segregation
- Dependency Inversion

### Design Patterns
- Repository Pattern
- Dependency Injection
- DTO Pattern
- Factory Pattern

### Clean Architecture
- Independência de Frameworks
- Testabilidade
- Independência de UI
- Independência de DB
- Dependency Rule

---

## 📚 Comandos Úteis

```bash
# Docker
docker-compose up -d              # Iniciar
docker-compose down               # Parar
docker-compose logs -f            # Ver logs

# Go
go mod tidy                       # Organizar dependências
go run cmd/clean_architecture_activity/main.go         # Executar
go build -o bin/server cmd/clean_architecture_activity/main.go  # Compilar

# Makefile
make help                         # Ver comandos
make install                      # Instalar
make run                          # Executar
make test                         # Testar
```

---

## 🎯 Checklist de Requisitos

### Funcionalidades
- [x] REST API - POST /order
- [x] REST API - GET /order
- [x] gRPC - CreateOrder
- [x] gRPC - ListOrders
- [x] GraphQL - createOrder
- [x] GraphQL - listOrders

### Infraestrutura
- [x] MongoDB com Docker
- [x] Dockerfile
- [x] docker-compose.yaml
- [x] Migrações (001_init.js)

### Arquitetura
- [x] Clean Architecture
- [x] Entities com validação
- [x] Use Cases com DTOs
- [x] Repository Pattern
- [x] Dependency Injection

### Documentação
- [x] README.md
- [x] api.http
- [x] Explicação das portas
- [x] Passos de execução

---

## 🚀 Início Rápido

```bash
# 1. Instalar
./scripts/install.sh

# 2. MongoDB e App
docker-compose up -d

# 4. Testar
curl http://localhost:8080/order
```

**Sistema rodando com REST (8080), gRPC (50051) e GraphQL (8081)!** 🎉