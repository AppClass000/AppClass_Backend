package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		JWTSecretKey := os.Getenv("JWTSECRETKEY")
		if JWTSecretKey == "" {
			log.Println("環境変数の取得に失敗しました")
		}
		Claims := &Claims{}

		tokenString, _ := c.Cookie("jwt")
		if tokenString == "" {
			log.Println("cookieがありません")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization Header missing"})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, Claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("署名方法が期待するものと違います %v", t.Header["alg"])
			}
			return []byte(JWTSecretKey), nil
		})
		if err != nil {
			fmt.Errorf("JWTパースエラー %v", err)
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "Invalid token",
					"isLoggin": false})
			c.Abort()
			return
		}
		if !token.Valid {
			fmt.Errorf("tokenが無効です")
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "Invalid token",
					"isLoggin": false,
				})
			c.Abort()
			return
		}

		c.Set("userID", Claims.UserID)
		value, exist := c.Get("userID")
		if exist {
			log.Println("useridです:", value)
		} else {
			log.Println("userID が Context にセットされていません")
		}
		c.Next()
	}
}
