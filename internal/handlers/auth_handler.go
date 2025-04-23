package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/0xDevvvvv/makerble/internal/models"
	"github.com/0xDevvvvv/makerble/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.BindJSON(&req)
	if req.Password == "" || req.Username == "" || err != nil {
		fmt.Println(req)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Please provide a username and password"})
		return
	}

	jwt, err := services.ValidateUser(h.DB, req.Username, req.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})

}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req models.UserCreate
	err := c.BindJSON(&req)
	if req.Password == "" || req.Username == "" || req.Role == "" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Please provide a username and password and a role"})
		return
	}
	user, err := services.CreateUser(h.DB, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "details": user})
}
