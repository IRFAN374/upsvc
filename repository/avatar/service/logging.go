package service

import (
	"context"
	"time"

	service "github.com/IRFAN374/upsvc/repository/avatar"
	log "github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Repository
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Repository) service.Repository {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (M loggingMiddleware) Add(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Add",
			"request", logAddRequest{},
			"response", logAddResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Add(arg0)
}

func (M loggingMiddleware) Update(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Update",
			"request", logUpdateRequest{},
			"response", logUpdateResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Update(arg0)
}

func (M loggingMiddleware) Remove(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Remove",
			"request", logRemoveRequest{},
			"response", logRemoveResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Remove(arg0)
}

func (M loggingMiddleware) GetById(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetById",
			"request", logGetByIdRequest{},
			"response", logGetByIdResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.GetById(arg0)
}

func (M loggingMiddleware) GetByName(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetByName",
			"request", logGetByNameRequest{},
			"response", logGetByNameResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.GetByName(arg0)
}

type (
	logAddRequest  struct{}
	logAddResponse struct{}

	logUpdateRequest  struct{}
	logUpdateResponse struct{}

	logRemoveRequest  struct{}
	logRemoveResponse struct{}

	logGetByIdRequest  struct{}
	logGetByIdResponse struct{}

	logGetByNameRequest  struct{}
	logGetByNameResponse struct{}
)
