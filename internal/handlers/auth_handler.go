package handlers

import (
	"net/http"

	"github.com/0xDevvvvv/makerble/internal/models"
	"github.com/0xDevvvvv/makerble/internal/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.BindJSON(&req)
	if req.Password == "" || req.Username == "" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Please provide a username and password"})
		return
	}

	jwt, err := services.ValidateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})

}
