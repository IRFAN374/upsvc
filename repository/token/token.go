package token

import "context"

type Repository interface {
	GenerateToken(ctx context.Context) (err error)
	UpdateToken(ctx context.Context) (err error)
	VerifiedToken(ctx context.Context) (err error)
}
