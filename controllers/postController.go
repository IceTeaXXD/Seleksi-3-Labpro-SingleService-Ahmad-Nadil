package controllers

import (
	"singleservice/initializers"
	model "singleservice/models"
	// "singleservice/auth"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gin-gonic/gin"
	// "fmt"
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
    // get username and password from request body
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    err := c.BindJSON(&body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // authenticate user
	Users := []model.User{}
	initializers.DB.Where("username = ? AND password = ?", body.Username, body.Password).Find(&Users)
	// check if username and password match
	if len(Users) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Username and password do not match",
			"data":    nil,
		})
		return
	}

    // generate JWT token
    // token, err := auth.GenerateToken(Users[0].ID)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{
    //         "status":  "error",
    //         "message": "Failed to generate token",
    //         "data":    nil,
    //     })
    //     return
    // }

    // return user data and token
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "User authenticated successfully",
        "data": gin.H{
            "user": gin.H{
                "username": Users[0].Username,
                "name":     Users[0].Name,
            },
            // "token": token,
        },
    })
	return;
}