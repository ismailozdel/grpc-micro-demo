package types

import (
	"context"

	"github.com/ismailozdel/micro/common/proto/invoice"
)

type InvoiceServiceClient interface {
	CreateInvoice(context.Context, *invoice.CreateInvoiceRequest) ([]*invoice.Invoice, error)
	GetInvoice(context.Context, string) (*invoice.GetInvoiceResponse, error)
	ListInvoices(ctx context.Context) (*invoice.ListInvoicesResponse, error)
}
