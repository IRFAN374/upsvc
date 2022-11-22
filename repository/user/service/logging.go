package service

import (
	"context"
	"time"

	models "github.com/IRFAN374/upsvc/models"
	service "github.com/IRFAN374/upsvc/repository/user"
	log "github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next  service.Repository
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Repository) service.Repository {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}

	}
}

func (M loggingMiddleware) Register(arg0 context.Context, arg1 models.UserInfo) (res0 string, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Register",
			"request", logRegisterRequest{},
			"response", logRegisterResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Register(arg0, arg1)
}

func (M loggingMiddleware) Login(arg0 context.Context, arg1 string, arg2 string) (res0 models.UserResponse, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Login",
			"request", logLoginRequest{},
			"response", logLoginResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Login(arg0, arg1, arg2)
}

func (M loggingMiddleware) GetUserById(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUserById",
			"request", logGetUserByIdRequest{},
			"response", logGetUserByIdResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.GetUserById(arg0, arg1)

}

func (M loggingMiddleware) GetUserByName(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUserByName",
			"request", logGetUserByNameRequest{},
			"response", logGetUserByNameResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.GetUserByName(arg0, arg1)

}

func (M loggingMiddleware) GetUserByEmail(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUserByEmail",
			"request", logGetUserByEmailRequest{},
			"response", logGetUserByEmailResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.GetUserByEmail(arg0, arg1)
}

func (M loggingMiddleware) UpdateUser(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateUser",
			"request", logUpdateUserRequest{},
			"response", logUpdateUserResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.UpdateUser(arg0, arg1)

}

func (M loggingMiddleware) DeleteUser(arg0 context.Context, arg1 string) (res0 bool, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "DeleteUser",
			"request", logDeleteUserRequest{},
			"response", logDeleteUserResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.DeleteUser(arg0, arg1)

}

func (M loggingMiddleware) IsExist(arg0 context.Context, arg1 string) (res0 bool, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "IsExist",
			"request", logIsExistRequest{},
			"response", logIsExistResponse{},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.IsExist(arg0, arg1)

}

type (
	logRegisterRequest  struct{}
	logRegisterResponse struct{}

	logLoginRequest  struct{}
	logLoginResponse struct{}

	logGetUserByIdRequest  struct{}
	logGetUserByIdResponse struct{}

	logGetUserByNameRequest  struct{}
	logGetUserByNameResponse struct{}

	logGetUserByEmailRequest  struct{}
	logGetUserByEmailResponse struct{}

	logUpdateUserRequest  struct{}
	logUpdateUserResponse struct{}

	logDeleteUserRequest  struct{}
	logDeleteUserResponse struct{}

	logIsExistRequest  struct{}
	logIsExistResponse struct{}
)
