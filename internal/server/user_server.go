package server

import (
	"context"

	"connectrpc.com/connect"

	userv1 "github.com/AI1411/connectpj/gen/proto/user/v1"
	"github.com/AI1411/connectpj/internal/infra/db"
)

type UserServer struct {
	dbClient db.Client
}

func NewUserServer(dbClient db.Client) *UserServer {
	return &UserServer{
		dbClient: dbClient,
	}
}

func (u *UserServer) GetUser(ctx context.Context, in *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.GetUserResponse], error) {
	if err := in.Msg.Validate(); err != nil {
		return nil, err
	}

	var user userv1.User
	if err := u.dbClient.Conn(ctx).Table("users").Where("id = ?", in.Msg.Id).First(&user).Error; err != nil {
		return nil, err
	}

	res := connect.NewResponse(&userv1.GetUserResponse{
		User: &userv1.User{
			Id:    user.GetId(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		},
	})
	return res, nil
}

func (u *UserServer) ListUsers(ctx context.Context, in *connect.Request[userv1.ListUsersRequest]) (*connect.Response[userv1.ListUsersResponse], error) {
	if err := in.Msg.Validate(); err != nil {
		return nil, err
	}

	var users []userv1.User
	if err := u.dbClient.Conn(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, err
	}

	res := make([]*userv1.User, len(users))
	for i := range users {
		user := &users[i]
		res[i] = &userv1.User{
			Id:    user.GetId(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		}
	}

	return connect.NewResponse(&userv1.ListUsersResponse{
		Users: res,
	}), nil
}

func (u *UserServer) CreateUser(ctx context.Context, in *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.CreateUserResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) UpdateUser(ctx context.Context, in *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.UpdateUserResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) DeleteUser(ctx context.Context, in *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[userv1.DeleteUserResponse], error) {
	//TODO implement me
	panic("implement me")
}
