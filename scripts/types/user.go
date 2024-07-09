package types

import "time"

type ReqUser struct {
	Limit string
	Order string
}

type ReqCreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
