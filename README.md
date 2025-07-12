# Challenge Clean Architecture

Este é um projeto de exemplo que implementa uma API de pedidos utilizando Clean Architecture, com suporte a múltiplos protocolos de comunicação.

## 🚀 Tecnologias

- Go 1.23+
- MySQL (banco de dados)
- gRPC
- GraphQL
- REST API
- RabbitMQ (para mensageria)
- Wire (para injeção de dependências)

## 📦 Serviços e Portas

A aplicação expõe os seguintes serviços nas respectivas portas:

| Serviço   | Porta | Descrição                               |
|-----------|-------|-----------------------------------------|
| HTTP REST | 8000  | API REST tradicional                    |
| gRPC      | 50051 | API gRPC para comunicação entre serviços|
| GraphQL   | 8080  | Interface GraphQL                       |
| Playground| 8080  | Interface interativa do GraphQL         |

## 🛠️ Configuração do Ambiente

1. **Pré-requisitos**
   - Go 1.23 ou superior
   - Git
   - MySQL
   - RabbitMQ (opcional, para mensageria)

2. **Clonar o repositório**
   ```bash
   git clone https://github.com/victorgiudicissi/challenge-clean-architecture.git
   cd challenge-clean-architecture
   ```

## 🚀 Executando a aplicação

1. **Instalar dependências**
   ```bash
   go mod tidy
   ```

2. **Iniciar os serviços**
   ```bash
   make up
   ```

3. **Executar a aplicação**
   ```bash
   make run
   ```

## 📚 Endpoints

### REST API
- `GET /order` - Lista todos os pedidos

```bash
curl http://localhost:8000/order
```

- `POST /order` - Cria um novo pedido

```bash
curl -X POST http://localhost:8000/order -H "Content-Type: application/json" -d '{"id":"123e4567-e89b-12d3-a456-426614174000","Price":100.00,"Tax":10.00}'
```

### GraphQL
- **Playground**: http://localhost:8080

#### Exemplo de Mutação (Criar Pedido)
```graphql
mutation {
  createOrder(input: {
    id: "123e4567-e89b-12d3-a456-426614174000"
    Price: 100.00
    Tax: 10.00
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Exemplo de Consulta (Listar Pedidos)
```graphql
query {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC
Os serviços gRPC podem ser acessados na porta 50051. Use um cliente gRPC (como BloomRPC, grpcurl ou Postman) para interagir com os serviços.
