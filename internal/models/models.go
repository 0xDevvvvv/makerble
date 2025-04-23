package models

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Patient struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
type User struct {
	ID        int
	Username  string
	Role      string
	CreatedAt time.Time
}

type UserCreate struct {
	ID        int
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time
}
