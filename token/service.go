package token

import "context"

type Service interface {
	GenerateToken(ctx context.Context) (err error)
	UpdateToken(ctx context.Context) (err error)
	VerifiedToken(ctx context.Context) (err error)
}
