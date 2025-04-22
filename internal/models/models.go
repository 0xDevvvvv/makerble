package models

type LoginRequest struct {
	Username string `json:"user"`
	Password string `json:"password"`
}
