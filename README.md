# 🚀 Microservices com gRPC - Sistema de E-commerce

Projeto acadêmico implementando uma arquitetura de microservices com comunicação gRPC para sistema de e-commerce, desenvolvido em Go com Docker Compose.

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Arquitetura](#arquitetura)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Pré-requisitos](#pré-requisitos)
- [Instalação e Execução](#instalação-e-execução)
- [Uso da API](#uso-da-api)
- [Testes](#testes)
- [Banco de Dados](#banco-de-dados)
- [Monitoramento](#monitoramento)
- [Troubleshooting](#troubleshooting)
- [Contribuição](#contribuição)

## 🎯 Sobre o Projeto

Este projeto implementa um sistema de e-commerce distribuído usando microservices com as seguintes funcionalidades:

- **Gestão de Pedidos**: Criação e gerenciamento de pedidos de clientes
- **Processamento de Pagamentos**: Validação e processamento de transações
- **Gestão de Entrega**: Cálculo de prazos de entrega e logística
- **Comunicação gRPC**: Comunicação eficiente entre serviços
- **Banco de Dados**: Persistência de dados com MySQL

### ✨ Funcionalidades Principais

- 🛒 Criação de pedidos com múltiplos produtos
- 💳 Validação de pagamentos com regras de negócio
- 📦 Cálculo automático de prazos de entrega
- 🔄 Comunicação síncrona entre microservices via gRPC
- 🐳 Containerização completa com Docker Compose
- 📊 Banco de dados relacional com múltiplas databases

## 🏗️ Arquitetura

O projeto segue os princípios da **Arquitetura Hexagonal** (Ports & Adapters) e está organizado em três microservices principais:

```
┌─────────────────┐    gRPC     ┌─────────────────┐    gRPC     ┌─────────────────┐
│                 │ ──────────► │                 │ ──────────► │                 │
│  Order Service  │             │ Payment Service │             │Shipping Service │
│     :3000       │             │      :3001      │             │      :3002      │
│                 │             │                 │             │                 │
└─────────┬───────┘             └─────────┬───────┘             └─────────┬───────┘
          │                               │                               │
          │                               │                               │
          └───────────────────────────────┼───────────────────────────────┘
                                          │
                                          ▼
                                ┌─────────────────┐
                                │                 │
                                │  MySQL 8.0      │
                                │    :3306        │
                                │                 │
                                └─────────────────┘
```

### 🔧 Microservices

#### 1. **Order Service** (Porta 3000)
- Gerencia criação e validação de pedidos
- Coordena comunicação com Payment e Shipping
- Valida produtos e quantidades
- Calcula totais dos pedidos

#### 2. **Payment Service** (Porta 3001)
- Processa validações de pagamento
- Aplica regras de negócio (limite R$ 1000, max 50 unidades)
- Gera registros de transações
- Simula processamento de pagamento

#### 3. **Shipping Service** (Porta 3002)
- Calcula prazos de entrega
- Fórmula: `delivery_days = 1 + (total_quantity / 5)`
- Gera códigos de rastreamento
- Gerencia logística de envio

## 🛠️ Tecnologias Utilizadas

- **Backend**: Go (Golang) 1.21+
- **Comunicação**: gRPC + Protocol Buffers
- **Banco de Dados**: MySQL 8.0
- **Containerização**: Docker + Docker Compose
- **Arquitetura**: Hexagonal (Ports & Adapters)
- **ORM**: GORM (Go Object Relational Mapping)

## Funcionalidades Implementadas

### Microsserviço Order
- ✅ Validação de produtos existentes no banco de dados
- ✅ Integração com Payment (chamada somente se produtos válidos)
- ✅ Integração com Shipping (chamada somente se pagamento bem-sucedido)
- ✅ Tratamento de erros apropriado
- ✅ Arquitetura hexagonal

### Microsserviço Shipping
- ✅ Recebe itens e ID da compra
- ✅ Calcula prazo de entrega (1 dia + 1 dia a cada 5 unidades)
- ✅ Arquitetura hexagonal
- ✅ Persistência em banco de dados

### Microsserviço Payment
- ✅ Processa pagamentos
- ✅ Arquitetura hexagonal existente

## Executando com Docker Compose

### Pré-requisitos
- Docker
- Docker Compose

### Passos

1. **Clone o repositório e navegue até o diretório:**
   ```bash
   cd microservices
   ```

2. **Execute os containers:**
   ```bash
   docker compose up --build -d
   ```

3. **Aguarde os serviços iniciarem:**
   - MySQL: porta 3306
   - Order: porta 3000
   - Payment: porta 3001
   - Shipping: porta 3002

4. **Verifique se os serviços estão rodando:**
   ```bash
   docker compose ps
   ```

## Testando a API

### Exemplo de requisição (gRPC)

Você pode usar uma ferramenta como `grpcurl` para testar:

```bash
# Testar criação de pedido
grpcurl -plaintext -d '{
  "customer_id": 1,
  "order_items": [
    {
      "product_code": "PROD001",
      "unit_price": 10.50,
      "quantity": 3
    },
    {
      "product_code": "PROD002", 
      "unit_price": 25.00,
      "quantity": 2
    }
  ]
}' localhost:3000 Order/Create
```

## Banco de Dados

O projeto inicializa automaticamente com:

- **Database `order`**: tabelas de pedidos e produtos
- **Database `payment`**: tabela de pagamentos  
- **Database `shipping`**: tabelas de envio
- **Produtos de exemplo**: PROD001 a PROD005

## Estrutura do Projeto

```
microservices/
├── order/              # Microsserviço de pedidos
├── payment/            # Microsserviço de pagamentos
├── shipping/           # Microsserviço de envio
├── docker-compose.yml  # Configuração Docker Compose
└── db.sql             # Script de inicialização do banco

microservices-proto/    # Definições Protocol Buffers
├── order/
├── payment/ 
└── shipping/
```

## Desenvolvimento

### Executando localmente

1. **Inicie o MySQL:**
   ```bash
   docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=minhasenha -p 3306:3306 mysql:8.0
   ```

2. **Configure as variáveis de ambiente:**
   ```bash
   export DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:3306)/order"
   export APPLICATION_PORT=3000
   export ENV=development
   export PAYMENT_SERVICE_URL=localhost:3001
   export SHIPPING_SERVICE_URL=localhost:3002
   ```

3. **Execute cada serviço:**
   ```bash
   # Terminal 1 - Payment
   cd payment && go run cmd/main.go
   
   # Terminal 2 - Shipping  
   cd shipping && go run cmd/main.go
   
   # Terminal 3 - Order
   cd order && go run cmd/main.go
   ```

## Logs e Troubleshooting

- **Verificar logs dos containers:**
  ```bash
  docker compose logs order
  docker compose logs payment
  docker compose logs shipping
  docker compose logs mysql
  ```

- **Acessar container para debug:**
  ```bash
  docker compose exec order sh
  ```

## Regras de Negócio

1. **Validação de Produtos**: Todos os produtos do pedido devem existir na tabela `products`
2. **Limite de Quantidade**: Máximo 50 unidades por pedido
3. **Fluxo de Processo**: Order → Payment → Shipping
4. **Cálculo de Entrega**: 1 dia + 1 dia a cada 5 unidades adicionais
5. **Tratamento de Erro**: Se payment falha, order é cancelado. Se shipping falha, apenas log é gerado.

## Pontos Extras (Bônus)

- ✅ **Docker**: Projeto totalmente containerizado
- ✅ **Documentação**: README completo com instruções
- ✅ **Arquitetura Hexagonal**: Implementada em todos os serviços
- ✅ **Tratamento de Erros**: Códigos gRPC apropriados
