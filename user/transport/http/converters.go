package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/IRFAN374/upsvc/common/chttp"
	"github.com/IRFAN374/upsvc/user/transport"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)

	return nil

}

func CommonHTTPResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return chttp.EncodeResponse(ctx, w, response)
}

func Decode_Register_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_Login_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_GetUserById_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetUserByIdRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_GetUserByName_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetUserByNameRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_GetUserByEmail_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetUserByEmailRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_UpdateUser_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.UpdateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_DeleteUser_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.DeleteUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_IsExist_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.IsExistRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func Decode_Register_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.RegisterResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_Login_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.LoginResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_GetUserById_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetUserByIdResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_GetUserByName_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetUserByNameResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_GetUserByEmail_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetUserByEmailResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_UpdateUser_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.UpdateUserResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_DeleteUser_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.DeleteUserResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Decode_IsExist_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.IsExistResponse

	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Encode_Register_Request(ctx context.Context, r *http.Request, request interface{}) error {

	_ = request.(*transport.RegisterRequest)

	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/register"))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_Login_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.LoginRequest)

	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/login"))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_GetUserById_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.GetUserByIdRequest)

	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/getUserById/{%s}", req.UserId))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_GetUserByName_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.GetUserByNameRequest)

	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/getUserByName/{%s}", req.UserName))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_GetUserByEmail_Request(ctx context.Context, r *http.Request, request interface{}) error {

	req := request.(*transport.GetUserByEmailRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/getUserByEmail/{%s}", req.Email))

	return CommonHTTPRequestEncoder(ctx, r, request)

}

func Encode_UpdateUser_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.UpdateUserRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/updateUser/{%s}", req.UserId))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_DeleteUser_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.DeleteUserRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/deleteUser/{%s}", req.UserId))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_IsExist_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.IsExistRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/isExists/{%s}", req.UserId))

	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_Register_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_Login_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_GetUserById_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_GetUserByName_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_GetUserByEmail_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_UpdateUser_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_DeleteUser_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func Encode_IsExist_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
