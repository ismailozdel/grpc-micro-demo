package stock

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ismailozdel/micro/common/proto/stock"
	"github.com/ismailozdel/micro/gateway/global"
	"google.golang.org/grpc"
)

type StockHandler struct {
	conn stock.StockServiceClient
}

func NewStockHandler(conn *grpc.ClientConn) *StockHandler {
	return &StockHandler{conn: stock.NewStockServiceClient(conn)}
}

func (h *StockHandler) GetStocks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := h.conn.GetStocks(ctx, &stock.GetStocksRequest{})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetStocks())
}

func (h *StockHandler) CreateStock(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := new(stock.Stock)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.conn.CreateStock(ctx, &stock.CreateStockRequest{Name: req.Name, Symbol: req.Symbol, Price: req.Price})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetStock())
}
