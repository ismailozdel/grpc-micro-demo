```

### 2. Generate Go code from proto file

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative invoice.proto
```

### 3. Implement the service

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "time"

    "github.com/ismailozdel/micro/proto/invoice"
    "google.golang.org/grpc"
)

type server struct {
    invoice.UnimplementedInvoiceServiceServer
}


func (s *server) CreateInvoice(ctx context.Context, req *invoice.CreateInvoiceRequest) (*invoice.CreateInvoiceResponse, error) {
    fmt.Println("CreateInvoice")
    return &invoice.CreateInvoiceResponse{
        Id: "1",
    }, nil
}

func (s *server) GetInvoice(ctx context.Context, req *invoice.GetInvoiceRequest) (*invoice.GetInvoiceResponse, error) {
    fmt.Println("GetInvoice")
    return &invoice.GetInvoiceResponse{
        Id: "1",
        Name: "Invoice 1",
        Description: "Description 1",
        Amount: 100,
    }, nil
}


func (s *server) ListInvoices(ctx context.Context, req *invoice.ListInvoicesRequest) (*invoice.ListInvoicesResponse, error) {
    fmt.Println("ListInvoices")
    return &invoice.ListInvoicesResponse{
        Invoices: []*invoice.ListInvoicesResponse_Invoice{
            {
                Id: "1",
                Name: "Invoice 1",
                Description: "Description 1",
                Amount: 100,
            },
            {
                Id: "2",
                Name: "Invoice 2",
                Description: "Description 2",
                Amount: 200,
            },
        },
    }, nil
}


func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    invoice.RegisterInvoiceServiceServer(s, &server{})

    fmt.Println("Starting server on port :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}


```

### 4. Run the server

```bash

go run main.go


```

### 5. Test the server

```bash

grpcurl -plaintext -d '{"name": "Invoice 1", "description": "Description 1", "amount": 100}' localhost:50051 invoice.InvoiceService/CreateInvoice

grpcurl -plaintext -d '{"id": "1"}' localhost:50051 invoice.InvoiceService/GetInvoice

grpcurl -plaintext localhost:50051 invoice.InvoiceService/ListInvoices

```

### 6. Implement the client

```go

package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ismailozdel/micro/proto/invoice"
    "google.golang.org/grpc"
)


func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := invoice.NewInvoiceServiceClient(conn)

    createInvoice(c)
    getInvoice(c)
    listInvoices(c)
}


func createInvoice(c invoice.InvoiceServiceClient) {
    fmt.Println("CreateInvoice")
    res, err := c.CreateInvoice(context.Background(), &invoice.CreateInvoiceRequest{
        Name: "Invoice 1",
        Description: "Description 1",
        Amount: 100,
    })
    if err != nil {
        log.Fatalf("CreateInvoice failed: %v", err)
    }
    fmt.Println(res)
}


func getInvoice(c invoice.InvoiceServiceClient) {
    fmt.Println("GetInvoice")
    res, err := c.GetInvoice(context.Background(), &invoice.GetInvoiceRequest{
        Id: "1",
    })
    if err != nil {
        log.Fatalf("GetInvoice failed: %v", err)
    }
    fmt.Println(res)
}


func listInvoices(c invoice.InvoiceServiceClient) {
    fmt.Println("ListInvoices")
    res, err := c.ListInvoices(context.Background(), &invoice.ListInvoicesRequest{})
    if err != nil {
        log.Fatalf("ListInvoices failed: %v", err)
    }
    fmt.Println(res)
}

```

### 7. Run the client

```bash

go run main.go

```

### 8. Run the server in Docker

```bash

docker build -t invoice-service .

docker run -p 50051:50051 invoice-service


```

### 9. Run the client in Docker

```bash

docker build -t invoice-client -f Dockerfile.client .

docker run invoice-client

```

### 10. Run the server in Kubernetes

```bash

kubectl apply -f deployment.yaml

kubectl apply -f service.yaml

```

### 11. Run the client in Kubernetes

```bash

kubectl apply -f job.yaml

```

### 12. Clean up

```bash

kubectl delete -f deployment.yaml

kubectl delete -f service.yaml

kubectl delete -f job.yaml

```

### 13. Clean up Docker

```bash

docker stop $(docker ps -a -q)

docker rm $(docker ps -a -q)

docker rmi $(docker images -a -q)

```
