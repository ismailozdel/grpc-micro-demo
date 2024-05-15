package stock

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func Router(router fiber.Router, conn *grpc.ClientConn) {
	h := NewStockHandler(conn)

	router.Post("/stock", h.CreateStock)
	router.Get("/stock", h.GetStocks)
}
