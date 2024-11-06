# 🛠️ Microservices System in Go (RESTful API & gRPC)

This project demonstrates a simple microservices architecture in Go, using `gin` for RESTful APIs and `gRPC` for inter-service communication. The system includes a `UserService` for managing user data and an `OrderService` for managing orders and retrieving user information.

## Project Overview

### Microservices

1. **UserService**: 
   - Exposes a gRPC service to manage user data.
   - Listens on port `50051`.

2. **OrderService**: 
   - Provides a RESTful API using `gin` to manage orders.
   - Uses gRPC to communicate with `UserService` for fetching user details.
   - Listens on port `8080`.

## Technologies Used

- **Go**: The programming language used for both microservices.
- **gRPC**: For efficient communication between `UserService` and `OrderService`.
- **Protocol Buffers (protobuf)**: For defining gRPC service and message formats.
- **gin**: A high-performance HTTP web framework for building RESTful APIs.

## Setup Instructions

### Prerequisites

1. **Go**: Make sure Go is installed on your system. [Download Go](https://golang.org/dl/)
2. **Protocol Buffers Compiler (protoc)**: Install `protoc` from [Protocol Buffers releases](https://github.com/protocolbuffers/protobuf/releases).
3. **Install `protoc-gen-go` and `protoc-gen-go-grpc`**:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
