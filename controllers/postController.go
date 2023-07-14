package controllers

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		User  struct {
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"user"`
		Token string `json:"token"`
	} `json:"data"`
}

func generateToken(username string) string {
	// Create the claims containing user information
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte("your-secret-key") // Replace with your own secret key
	tokenString, _ := token.SignedString(secretKey)

	return tokenString
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request"})
		return
	}

	// Authenticate user
	if req.Username != "admin" || req.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid credentials"})
		return
	}
	
	// Generate JWT token
	token := generateToken(req.Username)

	// Prepare response
	resp := LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Data: struct {
			User  struct {
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"user"`
			Token string `json:"token"`
		}{
			User: struct {
				Username string `json:"username"`
				Name     string `json:"name"`
			}{
				Username: req.Username,
				Name:     "Administrator",
			},
			Token: token,
		},
	}

	c.JSON(http.StatusOK, resp)
}