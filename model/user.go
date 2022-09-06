package model

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tructn/redstring/env"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var secretKey []byte

func init() {
	env := env.GetEnv()
	secretKey = []byte(env.Keys.SecretKey)
}

type User struct {
	gorm.Model
	UserName     string `json:"userName"`
	HashPassword string `json:"hashPassword"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	IsActive     bool   `json:"isActive"`
}

func (u *User) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.HashPassword = string(bytes)
	return string(bytes), err
}

func (u User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashPassword), []byte(password))
	return err == nil
}

func (u User) GenerateToken() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		Issuer:    "redstring",
		Subject:   u.UserName,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)

	fmt.Println(string(secretKey))

	tokenStr, err := token.SignedString(string(secretKey))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
