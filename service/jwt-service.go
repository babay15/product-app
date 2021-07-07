package service

import "github.com/dgrijalva/jwt-go"

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
}