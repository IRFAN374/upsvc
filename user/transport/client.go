package transport

import (
	"context"

	models "github.com/IRFAN374/upsvc/models"
)

func (set EndpointsSet) Register(arg0 context.Context, arg1 models.UserInfo) (res0 string, res1 error) {

	request := RegisterRequest{
		UserInfo: arg1,
	}
	response, res1 := set.RegisterEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*RegisterResponse).UserId, res1
}

func (set EndpointsSet) Login(arg0 context.Context, arg1 string, arg2 string) (res0 models.UserResponse, res1 error) {

	request := LoginRequest{
		Email:    arg1,
		Password: arg2,
	}
	response, res1 := set.LoginEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*LoginResponse).UserRes, res1
}

func (set EndpointsSet) GetUserById(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	request := GetUserByIdRequest{
		UserId: arg1,
	}
	response, res1 := set.GetUserByIdEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*GetUserByIdResponse).UserRes, res1
}

func (set EndpointsSet) GetUserByName(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	request := GetUserByNameRequest{
		UserName: arg1,
	}

	response, res1 := set.GetUserByNameEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*GetUserByNameResponse).UserRes, res1

}

func (set EndpointsSet) GetUserByEmail(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	request := GetUserByEmailRequest{
		Email: arg1,
	}

	response, res1 := set.GetUserByEmailEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*GetUserByEmailResponse).UserRes, res1

}

func (set EndpointsSet) UpdateUser(arg0 context.Context, arg1 string) (res0 models.UserResponse, res1 error) {

	request := UpdateUserRequest{
		UserId: arg1,
	}

	response, res1 := set.UpdateUserEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*UpdateUserResponse).UserRes, res1

}

func (set EndpointsSet) DeleteUser(arg0 context.Context, arg1 string) (res0 bool, res1 error) {

	request := DeleteUserRequest{
		UserId: arg1,
	}

	response, res1 := set.DeleteUserEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*DeleteUserResponse).Ok, res1
}

func (set EndpointsSet) IsExist(arg0 context.Context, arg1 string) (res0 bool, res1 error) {

	request := IsExistRequest{
		UserId: arg1,
	}

	response, res1 := set.IsExistEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*IsExistResponse).Ok, res1

}
