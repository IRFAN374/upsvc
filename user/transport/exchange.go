package transport

import models "github.com/IRFAN374/upsvc/models"

type (
	RegisterRequest struct {
		UserInfo models.UserInfo `json:"userInfo"`
	}
	RegisterResponse struct {
		UserId string `json:"userId"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginResponse struct {
		UserRes models.UserResponse `json:"userRes"`
	}

	GetUserByIdRequest struct {
		UserId string `json:"userId"`
	}
	GetUserByIdResponse struct {
		UserRes models.UserResponse `json:"userRes"`
	}

	GetUserByNameRequest struct {
		UserName string `json:"userName"`
	}
	GetUserByNameResponse struct {
		UserRes models.UserResponse `json:"userRes"`
	}

	GetUserByEmailRequest struct {
		Email string `json:"email"`
	}
	GetUserByEmailResponse struct {
		UserRes models.UserResponse `json:"userRes"`
	}

	UpdateUserRequest struct {
		UserId string `json:"userId"`
	}
	UpdateUserResponse struct {
		UserRes models.UserResponse `json:"userRes"`
	}

	DeleteUserRequest struct {
		UserId string `json:"userId"`
	}
	DeleteUserResponse struct {
		Ok bool `json:"ok"`
	}

	IsExistRequest struct {
		UserId string `json:"userId"`
	}
	IsExistResponse struct {
		Ok bool `json:"ok"`
	}
)
