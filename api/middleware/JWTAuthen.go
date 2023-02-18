package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte

func JWTAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signature ", token.Header["alg"])
			}

			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("userIDDDD", claims["userId"])
			// Set example variable
			c.Set("userId", claims["userId"])

		} else {

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":      "ok",
				"error":       err.Error(),
				"tokenString": tokenString,
				"mesage":      "Read user Fail",
			})

		}

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
