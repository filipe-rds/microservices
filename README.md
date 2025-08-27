# 🚀 Microservices com gRPC - Sistema de E-commerce

Projeto acadêmico implementando uma arquitetura de microservices com comunicação gRPC para sistema de e-commerce, desenvolvido em Go com Docker Compose.

## 📋 Índice

- [Sobre o Projeto](#-sobre-o-projeto)
- [Arquitetura](#️-arquitetura)
- [Tecnologias Utilizadas](#️-tecnologias-utilizadas)
- [Funcionalidades Implementadas](#funcionalidades-implementadas)
- [Executando com Docker Compose](#executando-com-docker-compose)
- [Testando a API](#testando-a-api)
- [Banco de Dados](#banco-de-dados)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Desenvolvimento](#desenvolvimento)
- [Logs e Troubleshooting](#logs-e-troubleshooting)
- [Regras de Negócio](#regras-de-negócio)
- [Pontos Extras](#pontos-extras-bônus)

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
   docker compose build && docker compose up -d
   ```

3. **Aguarde os serviços iniciarem:**
   - Database: porta 3306
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

### Testes Essenciais

Para uma validação completa do sistema, execute os seguintes testes:

```bash
# Teste 1: Pedido válido simples
grpcurl -d '{"customer_id": 123, "order_items":[{"product_code": "PROD001", "quantity": 4, "unit_price": 10.50}]}' -plaintext localhost:3000 Order/Create

# Teste 2: Pedido com múltiplos produtos
grpcurl -d '{"customer_id": 456, "order_items":[{"product_code": "PROD001", "quantity": 3, "unit_price": 10.50}, {"product_code": "PROD002", "quantity": 2, "unit_price": 25.00}]}' -plaintext localhost:3000 Order/Create

# Teste 3: Erro - Valor acima do limite (deve falhar)
grpcurl -d '{"customer_id": 123, "order_items":[{"product_code": "PROD001", "quantity": 10, "unit_price": 120}]}' -plaintext localhost:3000 Order/Create

# Teste 4: Shipping Service - Cálculo de prazo
grpcurl -d '{"order_id": 100, "items":[{"product_code": "PROD001", "quantity": 10}]}' -plaintext localhost:3002 Shipping/Create
```

> **Dica**: Consulte o arquivo `comands.txt` para uma lista completa de comandos de teste.

## Banco de Dados

O projeto inicializa automaticamente com MySQL 8.0 e cria:

- **Database `order`**: tabelas orders, order_items, products
- **Database `payment`**: tabela payments  
- **Database `shipping`**: tabelas shippings, shipping_items
- **Produtos pré-cadastrados**: PROD001 a PROD005 com preços e estoque definidos

### Produtos Disponíveis:
- **PROD001**: Produto Premium 1 - R$ 10,50 (100 unidades)
- **PROD002**: Produto Especial 2 - R$ 25,00 (50 unidades)
- **PROD003**: Produto Padrão 3 - R$ 15,75 (75 unidades)
- **PROD004**: Produto Deluxe 4 - R$ 30,00 (25 unidades)
- **PROD005**: Produto Econômico 5 - R$ 12,99 (200 unidades)

### Acesso ao Banco:
```bash
# Acessar MySQL via Docker Compose
docker compose exec database mysql -uroot -pminhasenha

# Verificar databases criadas
docker compose exec database mysql -uroot -pminhasenha -e "SHOW DATABASES;"

# Ver produtos cadastrados
docker compose exec database mysql -uroot -pminhasenha -e "USE \`order\`; SELECT * FROM products;"
```

## Estrutura do Projeto

```
microservices/
├── order/              # Microservice de pedidos
├── payment/            # Microservice de pagamentos
├── shipping/           # Microservice de envio
├── database/           # Configurações e scripts de banco
│   └── init/
│       └── db.sql     # Script de inicialização do banco
├── docker-compose.yml  # Configuração Docker Compose
└── comands.txt        # Comandos essenciais para testes

microservices-proto/    # Definições Protocol Buffers
├── order/
├── payment/ 
└── shipping/
```

## Desenvolvimento

### Executando localmente

1. **Inicie o Database:**
   ```bash
   docker run -d --name database -e MYSQL_ROOT_PASSWORD=minhasenha -p 3306:3306 mysql:8.0
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
  docker compose logs database
  ```

- **Acessar container para debug:**
  ```bash
  docker compose exec order sh
  ```

## Regras de Negócio

1. **Validação de Produtos**: Todos os produtos do pedido devem existir na tabela `products`
2. **Limite de Valor**: Máximo R$ 1000 por pedido
3. **Limite de Quantidade**: Máximo 50 unidades por pedido
4. **Fluxo de Processo**: Order → Payment → Shipping
5. **Cálculo de Entrega**: 1 dia base + 1 dia adicional a cada 5 unidades
6. **Tratamento de Erro**: Se payment falha, order é cancelado. Se shipping falha, apenas log é gerado.

## Pontos Extras (Bônus)

- ✅ **Docker Compose**: Orquestração completa com container `microservices-database`
- ✅ **Arquitetura Hexagonal**: Implementada em todos os microservices
- ✅ **Tratamento de Erros**: Códigos gRPC apropriados para cada situação
- ✅ **Banco Organizado**: Estrutura `database/init/` para scripts de inicialização
- ✅ **Documentação Completa**: README detalhado + arquivo `comands.txt` com testes
- ✅ **Validações de Negócio**: Limite de valor R$ 1000 e quantidade 50 unidades
- ✅ **Cálculo de Entrega**: Algoritmo baseado em quantidade de itens
- ✅ **Comunicação gRPC**: Integração perfeita entre Order → Payment → Shipping
