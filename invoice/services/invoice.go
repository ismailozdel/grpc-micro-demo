package services

import (
	"context"
	"fmt"

	"github.com/ismailozdel/micro/common/proto/invoice"
	"github.com/ismailozdel/micro/common/proto/stock"
	"github.com/ismailozdel/micro/common/proto/user"
	"google.golang.org/grpc"
)

/*
type InvoiceServiceClient interface {
	CreateInvoice(context.Context, *invoice.Invoice) (*[]invoice.Invoice, error)
	GetInvoice(context.Context, string) (*[]invoice.Invoice, error)
	ListInvoices(ctx context.Context) (*[]invoice.Invoice, error)
}
*/

var invoiceList = []*invoice.Invoice{
	{
		Id:          "1",
		Name:        "John Doe",
		Description: "Test",
		Amount:      100,
		StockId:     "1",
		UserId:      "1",
	},
}

type InvoiceService struct {
	stockClient stock.StockServiceClient
	userClient  user.UserServiceClient
}

func NewInvoiceService(stockConn *grpc.ClientConn, userConn *grpc.ClientConn) *InvoiceService {
	return &InvoiceService{
		stockClient: stock.NewStockServiceClient(stockConn),
		userClient:  user.NewUserServiceClient(userConn),
	}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, i *invoice.CreateInvoiceRequest) ([]*invoice.Invoice, error) {

	id := fmt.Sprintf("%d", len(invoiceList)+1)
	invoiceList = append(invoiceList, &invoice.Invoice{
		Id:          id,
		Name:        i.Name,
		Description: i.Description,
		Amount:      i.Amount,
		StockId:     i.StockId,
		UserId:      i.UserId,
	})
	return invoiceList, nil

}

func (s *InvoiceService) GetInvoice(ctx context.Context, id string) (*invoice.GetInvoiceResponse, error) {
	var item *invoice.Invoice
	for _, i := range invoiceList {
		if i.Id == id {
			item = i
			resStock, err := s.stockClient.GetStock(ctx, &stock.GetStockRequest{Id: i.StockId})
			if err != nil {
				return nil, err
			}
			resUser, err := s.userClient.GetUser(ctx, &user.GetUserRequest{Id: i.UserId})
			if err != nil {
				return nil, err
			}
			r := invoice.GetInvoiceResponse{
				Invoice: []*invoice.InvoiceWithStockAndUser{
					{
						Invoice: item,
						Stock:   resStock.Stock,
						User:    resUser.User[0],
					},
				}}
			return &r, nil
		}
	}

	return nil, nil
}

func (s *InvoiceService) ListInvoices(ctx context.Context) (*invoice.ListInvoicesResponse, error) {
	r := invoice.ListInvoicesResponse{
		Invoices: []*invoice.InvoiceWithStockAndUser{},
	}
	for _, i := range invoiceList {
		resStock, err := s.stockClient.GetStock(ctx, &stock.GetStockRequest{Id: i.StockId})
		if err != nil {
			return nil, err
		}
		resUser, err := s.userClient.GetUser(ctx, &user.GetUserRequest{Id: i.UserId})
		if err != nil {
			return nil, err
		}
		var user *user.User
		if len(resUser.User) != 0 {
			user = resUser.User[0]
		}

		r.Invoices = append(r.Invoices, &invoice.InvoiceWithStockAndUser{
			Invoice: i,
			Stock:   resStock.Stock,
			User:    user,
		})
	}
	return &r, nil
}
