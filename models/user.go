package models

import "time"

type UserInfo struct {
	Id              int    `json:"id,omitempty"`
	UserId          string `json:"user_id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Passowrd        string `json:"passowrd,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`

	createdAt time.Time `json:"created_at,omitempty"`
	updatedAt time.Time `json:"updated_at,omitempty"`
}

type UserToken struct {
	UserId       string `json:"user_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ActiveToken  string `json:"active_token,omitempty"`

	createdAt time.Time `json:"created_at,omitempty"`
	updatedAt time.Time `json:"updated_at,omitempty"`
}

type UserAvatar struct {
	UserId   string `json:"user_id,omitempty"`
	AvatarId string `json:"avatar_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`

	createdAt time.Time `json:"created_at,omitempty"`
	updatedAt time.Time `json:"updated_at,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	Id           int    `json:"id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	Email        string `json:"email,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ActiveToken  string `json:"active_token,omitempty"`
	AvatarId     string `json:"avatar_id,omitempty"`
	ImageUrl     string `json:"image_url,omitempty"`
}
