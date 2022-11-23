package transporthttp

import (
	"net/url"

	transport "github.com/IRFAN374/upsvc/user/transport"
	httpkit "github.com/go-kit/kit/transport/http"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {

	return transport.EndpointsSet{
		RegisterEndpoint: httpkit.NewClient(
			"POST", u,
			Encode_Register_Request,
			Decode_Register_Response,
			opts...,
		).Endpoint(),
		LoginEndpoint: httpkit.NewClient(
			"POST", u,
			Encode_Login_Request,
			Decode_Login_Response,
			opts...,
		).Endpoint(),
		GetUserByIdEndpoint: httpkit.NewClient(
			"GET", u,
			Encode_GetUserById_Request,
			Decode_GetUserById_Response,
			opts...,
		).Endpoint(),
		GetUserByNameEndpoint: httpkit.NewClient(
			"GET", u,
			Encode_GetUserByName_Request,
			Decode_GetUserByName_Response,
			opts...,
		).Endpoint(),
		GetUserByEmailEndpoint: httpkit.NewClient(
			"GET", u,
			Encode_GetUserByEmail_Request,
			Decode_GetUserByEmail_Response,
			opts...,
		).Endpoint(),
		UpdateUserEndpoint: httpkit.NewClient(
			"PUT", u,
			Encode_UpdateUser_Request,
			Decode_UpdateUser_Response,
			opts...,
		).Endpoint(),
		DeleteUserEndpoint: httpkit.NewClient(
			"DELETE", u,
			Encode_DeleteUser_Request,
			Decode_DeleteUser_Response,
			opts...,
		).Endpoint(),
		IsExistEndpoint: httpkit.NewClient(
			"GET", u,
			Encode_IsExist_Request,
			Decode_IsExist_Response,
			opts...,
		).Endpoint(),
	}
}
