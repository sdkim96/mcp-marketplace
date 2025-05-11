package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sdkim96/mcp-marketplace/internal/config"
)

type JWTClaims struct {
	Sub string `json:"sub"`
	Nbf int64  `json:"nbf"`
	Exp int64  `json:"exp"`
}

func GetJWTtoken(
	secret string,
	userID string,
	duration int,
) string {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userID,
			"nbf": time.Now().Unix(),
			"exp": time.Now().Add(time.Duration(duration) * time.Second).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}
func validateJWTtoken(secret string, tokenString string) (string, bool) {

	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		return "유효하지 않습니다.", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "유효하지 않습니다.", false
	}

	// 만료 시간 확인
	exp, ok := claims["exp"].(float64)
	if !ok || int64(exp) < time.Now().Unix() {
		fmt.Println("Token expired")
		return "유효하지 않습니다.", false
	}

	// subject 추출
	sub, ok := claims["sub"].(string)
	if !ok {
		return "유효하지 않습니다.", false
	}

	return sub, true
}

func GlobalApplicationMiddleware(
	appConfig *config.AppConfig,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("projectName", appConfig.ProjectName())
		c.Set("projectSecret", appConfig.ProjectSecret())
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		secret, ok := c.Get("projectSecret")
		if !ok {
			log.Println("projectSecret not found in context")
			c.JSON(500, gin.H{"error": "projectSecret not found"})
			c.Abort()
			return
		}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		sub, ok := validateJWTtoken(secret.(string), token)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userName", sub)
		c.Next()
	}
}
