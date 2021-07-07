package middleware

import (
	"github.com/babay15/product-app/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func JwtAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":http.StatusUnauthorized,
				"message":"You are not authorized to access this service",
			})
		}
	}
}