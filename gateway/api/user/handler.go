package user

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ismailozdel/micro/common/proto/user"
	"github.com/ismailozdel/micro/gateway/global"
	"google.golang.org/grpc"
)

type UserHandler struct {
	conn user.UserServiceClient
}

func NewUserHandler(conn *grpc.ClientConn) *UserHandler {
	return &UserHandler{conn: user.NewUserServiceClient(conn)}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := h.conn.GetUser(ctx, &user.GetUserRequest{
		Id: c.Params("id"),
	})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetUser())
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := new(user.User)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.conn.CreateUser(ctx, &user.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetUser())
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := new(user.User)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.conn.UpdateUser(ctx, &user.UpdateUserRequest{
		User: &user.User{
			Id:       req.Id,
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		},
	})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetUser())
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.conn.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: c.Params("id"),
	})
	if err != nil {
		return err
	}

	return global.SendResponse(c, nil, res.GetId())
}

/*
   rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse) {}
   rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {}
*/
