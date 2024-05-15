package types

import (
	"context"

	"github.com/ismailozdel/micro/common/proto/user"
)

type UserService interface {
	GetUser(context.Context, string) ([]*user.User, error)
	CreateUser(context.Context, *user.User) ([]*user.User, error)
	UpdateUser(context.Context, *user.User) ([]*user.User, error)
	DeleteUser(context.Context, string) error
}
