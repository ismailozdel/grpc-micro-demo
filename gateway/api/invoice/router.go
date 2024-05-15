package invoice

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func Router(router fiber.Router, conn *grpc.ClientConn) {
	h := NewInvoiceHandler(conn)

	router.Get("/invoice/:id", h.GetById)
	router.Post("/invoice", h.Create)
	router.Get("/invoice", h.ListInvoice)
}
