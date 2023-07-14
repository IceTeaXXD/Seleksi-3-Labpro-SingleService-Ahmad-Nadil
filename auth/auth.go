package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userID uint) (string, error) {
    // create claims
    claims := jwt.MapClaims{}
    claims["userID"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    // create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // sign token
    secret := []byte("mysecretkey")
    tokenString, err := token.SignedString(secret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}