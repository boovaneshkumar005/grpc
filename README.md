``markdown
# E-commerce gRPC Project in Golang

This is a **simple e-commerce gRPC project** implemented in **Go**, demonstrating:

- gRPC server and client communication
- Protobuf message definitions
- Clean folder structure with abstraction
- Internal packages for business logic
- Go modules and package management

---

## Table of Contents

- [Overview](#overview)
- [Folder Structure](#folder-structure)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Running the Project](#running-the-project)
- [gRPC Concepts](#grpc-concepts)
- [Project Details](#project-details)
- [Extending the Project](#extending-the-project)

---

## Overview

This project demonstrates a simple **ProductService** for an e-commerce system.  
It includes:

- `GetProduct` – fetch a product by ID  
- `ListProducts` – list all available products  

We use **Protocol Buffers** (`.proto`) for defining messages and services, and **gRPC** for client-server communication.

---

## Folder Structure

```

ecommerce-grpc/
│
├── go.mod                       # Go module
├── go.sum
│
├── proto/                        # Protobuf definitions
│   └── product.proto
│
├── proto/productpb/              # Generated Go code from proto
│   ├── product.pb.go
│   └── product_grpc.pb.go
│
├── internal/                     # Internal packages (business logic)
│   ├── product/
│   │   ├── model.go              # Product struct
│   │   ├── repository.go         # Repository interface & implementation
│   │   └── service.go            # Business logic & gRPC handler
│   │
│   └── server/
│       └── grpc.go               # gRPC server setup
│
├── pkg/                          # Optional shared packages
│   └── logger/
│       └── logger.go
│
└── cmd/
├── server/
│   └── main.go               # gRPC server entrypoint
└── client/
└── main.go               # gRPC client

````

---

## Dependencies

- Go >= 1.21
- gRPC for Go: `google.golang.org/grpc`
- Protobuf Go plugin: `google.golang.org/protobuf/cmd/protoc-gen-go`
- gRPC Go plugin: `google.golang.org/grpc/cmd/protoc-gen-go-grpc`
- Protocol Buffers compiler (`protoc`)

---

## Setup

### 1. Clone the repository
```bash
git clone https://github.com/m048/ecommerce-grpc.git
cd ecommerce-grpc
````

### 2. Install Go plugins

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. Install Protocol Buffers compiler (`protoc`) on Windows

* Download from: [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
* Extract and add `bin` folder to PATH
* Test:

```powershell
protoc --version
```

### 4. Generate Go code from `.proto`

```powershell
protoc --go_out=. --go-grpc_out=. proto/product.proto
```

---

## Running the Project

### Run Server

```powershell
cd cmd\server
go run main.go
```

Output:

```
gRPC server running on :50051
```

### Run Client (in a new terminal)

```powershell
cd cmd\client
go run main.go
```

Example output:

```
👉 GetProduct Response:
ID: 1, Name: Laptop, Price: 1200.00

👉 ListProducts Response:
ID: 1, Name: Laptop, Price: 1200.00
ID: 2, Name: Phone, Price: 800.00
```

---

## gRPC Concepts

* **gRPC** is a high-performance RPC framework using **HTTP/2**.
* **Protocol Buffers (`.proto`)** define:

  * Messages (data structures)
  * Services (RPC endpoints)
* Server implements the service interface, client uses the generated stub to call methods.
* Benefits:

  * Strongly typed contracts
  * Language-agnostic
  * Efficient binary serialization
  * Built-in streaming support

---

## Project Details

* **ProductService**

  * `GetProduct(GetProductRequest) → Product`
  * `ListProducts(ListProductsRequest) → ListProductsResponse`
* **Internal Abstraction**

  * Repository interface defines data access
  * Service layer implements business logic
  * gRPC handlers call service layer
* **Clean Architecture**

  * `internal/` – business logic, not exposed outside module
  * `cmd/` – entrypoints for server and client
  * `proto/` – contract files and generated code

---

## Extending the Project

You can extend this project easily:

1. Add new services (e.g., `OrderService`, `UserService`) in `proto/`.
2. Implement repository & service logic in `internal/`.
3. Register new services in `internal/server/grpc.go`.
4. Regenerate `.pb.go` files after proto changes.
5. Update the client to call new services.

---

## References

* [gRPC official docs](https://grpc.io/docs/)
* [Protocol Buffers docs](https://developers.google.com/protocol-buffers)
* [Go gRPC guide](https://pkg.go.dev/google.golang.org/grpc)

---
