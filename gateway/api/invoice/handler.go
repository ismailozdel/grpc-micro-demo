package invoice

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ismailozdel/micro/common/proto/invoice"
	"github.com/ismailozdel/micro/gateway/global"
	"google.golang.org/grpc"
)

type InvoiceHandler struct {
	conn invoice.InvoiceServiceClient
}

func NewInvoiceHandler(conn *grpc.ClientConn) *InvoiceHandler {
	return &InvoiceHandler{conn: invoice.NewInvoiceServiceClient(conn)}
}

func (h *InvoiceHandler) GetById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.conn.GetInvoice(ctx, &invoice.GetInvoiceRequest{Id: c.Params("id")})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetInvoice())
}

func (h *InvoiceHandler) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := new(invoice.Invoice)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.conn.CreateInvoice(ctx, &invoice.CreateInvoiceRequest{Name: req.Name, Amount: req.Amount, Description: req.Description, StockId: req.StockId, UserId: req.UserId})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetInvoice())
}

func (h *InvoiceHandler) ListInvoice(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.conn.ListInvoices(ctx, &invoice.ListInvoicesRequest{})
	if err != nil {
		log.Println(err)
		return err
	}

	return global.SendResponse(c, nil, res.GetInvoices())
}
