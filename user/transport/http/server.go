package transporthttp

import (
	http "net/http"

	transport "github.com/IRFAN374/upsvc/user/transport"
	httpkit "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...httpkit.ServerOption) http.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/api/register").Handler(
		httpkit.NewServer(
			endpoints.RegisterEndpoint,
			Decode_Register_Request,
			Encode_Register_Response,
			opts...))
	mux.Methods("POST").Path("/api/login").Handler(
		httpkit.NewServer(
			endpoints.LoginEndpoint,
			Decode_Login_Request,
			Encode_Login_Response,
			opts...))
	mux.Methods("GET").Path("/api/getUserById").Handler(
		httpkit.NewServer(
			endpoints.GetUserByIdEndpoint,
			Decode_GetUserById_Request,
			Encode_GetUserById_Response,
			opts...))
	mux.Methods("GET").Path("/api/getUserByName").Handler(
		httpkit.NewServer(
			endpoints.GetUserByNameEndpoint,
			Decode_GetUserByName_Request,
			Encode_GetUserByName_Response,
			opts...))
	mux.Methods("GET").Path("/api/getUserByEmail").Handler(
		httpkit.NewServer(
			endpoints.GetUserByEmailEndpoint,
			Decode_GetUserByEmail_Request,
			Encode_GetUserByEmail_Response,
			opts...))
	mux.Methods("PUT").Path("/api/updateUser").Handler(
		httpkit.NewServer(
			endpoints.UpdateUserEndpoint,
			Decode_UpdateUser_Request,
			Encode_UpdateUser_Response,
			opts...))

	mux.Methods("DELETE").Path("/api/deleteUser").Handler(
		httpkit.NewServer(
			endpoints.DeleteUserEndpoint,
			Decode_DeleteUser_Request,
			Encode_DeleteUser_Response,
			opts...))
	mux.Methods("GET").Path("/api/isExists").Handler(
		httpkit.NewServer(
			endpoints.IsExistEndpoint,
			Decode_IsExist_Request,
			Encode_IsExist_Response,
			opts...))

	return mux
}
