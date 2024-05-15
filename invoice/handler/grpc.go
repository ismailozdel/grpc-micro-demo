package handler

import (
	"context"

	"github.com/ismailozdel/micro/common/proto/invoice"
	"github.com/ismailozdel/micro/invoice/types"
	"google.golang.org/grpc"
)

type InvoiceGrpcHandler struct {
	invoceService types.InvoiceServiceClient
	invoice.UnimplementedInvoiceServiceServer
}

func NewInvoiceGrpcHandler(grpc *grpc.Server, invoiceService types.InvoiceServiceClient) {
	grpcHandler := &InvoiceGrpcHandler{
		invoceService: invoiceService,
	}
	invoice.RegisterInvoiceServiceServer(grpc, grpcHandler)
}

func (h *InvoiceGrpcHandler) CreateInvoice(ctx context.Context, req *invoice.CreateInvoiceRequest) (*invoice.CreateInvoiceResponse, error) {
	items, err := h.invoceService.CreateInvoice(ctx, req)
	if err != nil {
		return nil, err
	}
	return &invoice.CreateInvoiceResponse{Invoice: items}, nil
}

func (h *InvoiceGrpcHandler) GetInvoice(ctx context.Context, req *invoice.GetInvoiceRequest) (*invoice.GetInvoiceResponse, error) {
	res, err := h.invoceService.GetInvoice(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *InvoiceGrpcHandler) ListInvoices(ctx context.Context, req *invoice.ListInvoicesRequest) (*invoice.ListInvoicesResponse, error) {
	res, err := h.invoceService.ListInvoices(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
