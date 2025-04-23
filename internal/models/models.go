package models

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Patient struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Address   string    `json:"address,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Illness   string    `json:"illness,omitempty"`
	CreatedAt time.Time `json:"created_at"`
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
