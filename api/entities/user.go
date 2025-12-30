package entities

import (
	"time"
)

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Address        	string `json:"address"`
	Phone        	string `json:"phone"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Gender     		int `json:"gender"`
	IsActive     	int `json:"is_active"`
}

type ResponseUser struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Address        	string `json:"address"`
	Phone        	string `json:"phone"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Gender     		int `json:"gender"`
	IsActive     	int `json:"is_active"`
}