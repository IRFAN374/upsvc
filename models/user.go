package models

import "time"

type UserInfo struct {
	Id        int
	UserId    string
	FirstName string
	LastName  string
	Email     string

	createdAt time.Time
	updatedAt time.Time
}

type UserToken struct {
	UserId       string
	RefreshToken string
	ActiveToken  string

	createdAt time.Time
	updatedAt time.Time
}

type UserAvatar struct {
	UserId   string
	AvatarId string
	ImageUrl string

	createdAt time.Time
	updatedAt time.Time
}
