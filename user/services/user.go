package services

import (
	"context"
	"fmt"

	"github.com/ismailozdel/micro/common/proto/user"
)

var userList = []*user.User{
	{
		Id:    "1",
		Name:  "John Doeee",
		Email: "iso",
	},
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context, id string) ([]*user.User, error) {

	for _, i := range userList {
		if i.Id == id {
			return []*user.User{i}, nil
		}
	}

	return nil, nil
}

func (s *UserService) CreateUser(ctx context.Context, u *user.User) ([]*user.User, error) {
	u.Id = fmt.Sprintf("%d", len(userList)+1)
	userList = append(userList, u)
	return userList, nil
}

func (s *UserService) UpdateUser(ctx context.Context, u *user.User) ([]*user.User, error) {

	for i, user := range userList {
		if user.Id == u.Id {
			userList[i] = u
		}
	}

	return userList, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	for i, user := range userList {
		if user.Id == id {
			fmt.Println("Deleted user: ", user)
			userList = append(userList[:i], userList[i+1:]...)
			fmt.Println("User list: ", userList)
			break
		}
	}

	return nil
}
