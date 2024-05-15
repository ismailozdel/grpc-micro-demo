package user

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func Router(router fiber.Router, conn *grpc.ClientConn) {
	h := NewUserHandler(conn)

	router.Get("/user/:id", h.GetUser)
	router.Post("/user", h.CreateUser)
	router.Put("/user", h.UpdateUser)
	router.Delete("/user/:id", h.DeleteUser)
}
