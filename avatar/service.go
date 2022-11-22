package avatar

import "context"

type Service interface {
	Add(ctx context.Context) (err error)
	Update(ctx context.Context) (err error)
	Remove(ctx context.Context) (err error)
	GetById(ctx context.Context) (err error)
	GetByName(ctx context.Context) (err error)
}
