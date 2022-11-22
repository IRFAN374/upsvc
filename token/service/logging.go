package service

import (
	"context"
	"time"

	service "github.com/IRFAN374/upsvc/token"
	log "github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (M loggingMiddleware) GenerateToken(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GenerateToken",
			"request", logGenerateTokenRequest{},
			"response", logGenerateTokenResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.GenerateToken(arg0)

}

func (M loggingMiddleware) UpdateToken(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateToken",
			"request", logUpdateTokenRequest{},
			"response", logUpdateTokenResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.UpdateToken(arg0)

}

func (M loggingMiddleware) VerifiedToken(arg0 context.Context) (res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "VerifiedToken",
			"request", logVerifiedTokenRequest{},
			"response", logVerifiedTokenResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.VerifiedToken(arg0)

}

type (
	logGenerateTokenRequest  struct{}
	logGenerateTokenResponse struct{}

	logUpdateTokenRequest  struct{}
	logUpdateTokenResponse struct{}

	logVerifiedTokenRequest  struct{}
	logVerifiedTokenResponse struct{}
)
