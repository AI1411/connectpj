package handler

import "context"

type UserHandler interface {
	GetUser(ctx context.Context) (*connect.Re, error)
	ListUsers(ctx context.Context) []string
	CreateUser(ctx context.Context) string
	UpdateUser(ctx context.Context) string
	DeleteUser(ctx context.Context) string
}
