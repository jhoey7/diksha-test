package utils

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// HashPassword is func to hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is function to check password hashing
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateToken is function for create token
func CreateToken(username string, secretKey []byte) (string, int64) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      exp,
		})

	tokenString, _ := token.SignedString(secretKey)
	return tokenString, exp
}
