package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ismailozdel/micro/gateway/api/invoice"
	"github.com/ismailozdel/micro/gateway/api/stock"
	"github.com/ismailozdel/micro/gateway/api/user"
	"google.golang.org/grpc"
)

func SetupRouter(app *fiber.App, invoiceGrpcConn *grpc.ClientConn, stockGrpcConn *grpc.ClientConn, userGrpcConn *grpc.ClientConn) {

	api := app.Group("/api")

	invoice.Router(api, invoiceGrpcConn)
	stock.Router(api, stockGrpcConn)
	user.Router(api, userGrpcConn)
}
