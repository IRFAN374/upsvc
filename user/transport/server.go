package transport

import (
	"context"

	user "github.com/IRFAN374/upsvc/user"
	endpoint "github.com/go-kit/kit/endpoint"
)

func Endpoints(svc user.Service) EndpointsSet {
	return EndpointsSet{
		RegisterEndpoint:       RegisterEndpoint(svc),
		LoginEndpoint:          LoginEndpoint(svc),
		GetUserByIdEndpoint:    GetUserByIdEndpoint(svc),
		GetUserByNameEndpoint:  GetUserByNameEndpoint(svc),
		GetUserByEmailEndpoint: GetUserByEmailEndpoint(svc),
		UpdateUserEndpoint:     UpdateUserEndpoint(svc),
		DeleteUserEndpoint:     DeleteUserEndpoint(svc),
		IsExistEndpoint:        IsExistEndpoint(svc),
	}
}

func RegisterEndpoint(svc user.Service) endpoint.Endpoint {

	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {

		req := request.(*RegisterRequest)
		res0, res1 := svc.Register(arg0, req.UserInfo)

		return &RegisterResponse{
			UserId: res0,
		}, res1
	}
}

func LoginEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*LoginRequest)

		res0, res1 := svc.Login(arg0, req.Email, req.Password)

		return &LoginResponse{
			UserRes: res0,
		}, res1
	}
}

func GetUserByIdEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetUserByIdRequest)

		res0, res1 := svc.GetUserById(arg0, req.UserId)

		return &GetUserByIdResponse{
			UserRes: res0,
		}, res1

	}

}

func GetUserByNameEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetUserByNameRequest)

		res0, res1 := svc.GetUserByName(arg0, req.UserName)

		return &GetUserByNameResponse{
			UserRes: res0,
		}, res1
	}
}

func GetUserByEmailEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetUserByEmailRequest)

		res0, res1 := svc.GetUserByEmail(arg0, req.Email)

		return &GetUserByEmailResponse{
			UserRes: res0,
		}, res1
	}
}

func UpdateUserEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdateUserRequest)

		res0, res1 := svc.UpdateUser(arg0, req.UserId)

		return &UpdateUserResponse{
			UserRes: res0,
		}, res1
	}
}

func DeleteUserEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*DeleteUserRequest)

		res0, res1 := svc.DeleteUser(arg0, req.UserId)

		return &DeleteUserResponse{
			Ok: res0,
		}, res1
	}
}

func IsExistEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*IsExistRequest)

		res0, res1 := svc.IsExist(arg0, req.UserId)

		return &IsExistResponse{
			Ok: res0,
		}, res1
	}
}
