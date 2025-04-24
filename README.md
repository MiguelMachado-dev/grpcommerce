# grpcommerce

A gRPC-based microservices architecture for an e-commerce platform, written in Go.

## Features

- **UserService**: Handles user registration, login, and profile management.
- **NotificationService**: Manages sending notifications to users. (TBD)
- **OrderService**: Processes customer orders. (TBD)
- **ProductService**: Manages the product catalog. (TBD)

## Prerequisites

- Go 1.24.2 or later
- PostgreSQL
- protocol buffers compiler (`protoc`)
- `protoc-gen-go` plugin
- `protoc-gen-go-grpc` plugin

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/MiguelMachado-dev/grpcommerce.git
   cd grpcommerce
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Generate gRPC code from protobuf definitions:

   ```bash
   protoc --go_out=. --go-grpc_out=. -I proto proto/*.proto
   ```

4. Set up PostgreSQL database:

   Ensure PostgreSQL is running and create the `ecommerce` database:

   ```bash
   psql -U postgres -c "CREATE DATABASE ecommerce;"
   ```

5. Configure environment variables (optional):

   Services will use the following defaults if not set:

   ```bash
   export PORT=50051
   export DATABASE_URL="postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable"
   ```

6. Run the services:

   ```bash
   go run services/user-service/main.go
   ```

   Once implemented, you can start other services similarly:

   ```bash
   go run services/notification-service/main.go
   go run services/order-service/main.go
   go run services/product-service/main.go
   ```

## Usage

Interact with UserService using `grpcurl`:

```bash
grpcurl -plaintext -d '{
  "email": "user@example.com",
  "username": "user",
  "password": "pass123",
  "firstName": "John",
  "lastName": "Doe"
}' localhost:50051 user.UserService/Register
```

Adjust the RPC method and payload for other calls.

## Project Structure

```
├── proto/                     # Protobuf service definitions
│   ├── user.proto
│   ├── notification.proto
│   ├── order.proto
│   └── product.proto
├── ecommerce/                 # Generated protobuf Go code
│   └── proto/
│       └── user/
│           ├── user.pb.go
│           └── user_grpc.pb.go
├── services/                  # gRPC service implementations
│   ├── user-service
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handler/
│   │   └── repository/
│   ├── notification-service
│   ├── order-service
│   └── product-service
├── go.mod
└── go.sum
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request to discuss changes.
