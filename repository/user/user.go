package user

import "context"

type Repository interface {
	Register(ctx context.Context) (err error)
	Login(ctx context.Context) (err error)
	GetUserById(ctx context.Context) (err error)
	GetUserByName(ctx context.Context) (err error)
	GetUserByEmail(ctx context.Context) (err error)

	UpdateUser(ctx context.Context) (err error)
	DeleteUser(ctx context.Context) (err error)
	IsExist(ctx context.Context) (oki bool, err error)
}
