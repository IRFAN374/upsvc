package user

import (
	"context"

	"github.com/IRFAN374/upsvc/models"
	"github.com/IRFAN374/upsvc/repository/avatar"
	"github.com/IRFAN374/upsvc/repository/token"
	"github.com/IRFAN374/upsvc/repository/user"
)

type Service interface {
	Register(ctx context.Context, userInfo models.UserInfo) (userId string, err error)
	Login(ctx context.Context, Email string, password string) (userRes models.UserResponse, err error)
	GetUserById(ctx context.Context, UserId string) (userRes models.UserResponse, err error)
	GetUserByName(ctx context.Context, UserName string) (userRes models.UserResponse, err error)
	GetUserByEmail(ctx context.Context, Email string) (userRes models.UserResponse, err error)

	UpdateUser(ctx context.Context, userId string) (userRes models.UserResponse, err error)
	DeleteUser(ctx context.Context, userId string) (ok bool, err error)
	IsExist(ctx context.Context, userId string) (ok bool, err error)
}

type service struct {
	userRepo   user.Repository
	avatarRepo avatar.Repository
	tokenRepo  token.Repository
}

func NewService(userRepo user.Repository, avatarRepo avatar.Repository, tokenRepo token.Repository) *service {

	return &service{
		userRepo:   userRepo,
		avatarRepo: avatarRepo,
		tokenRepo:  tokenRepo,
	}
}
