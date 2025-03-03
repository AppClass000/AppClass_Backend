package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

func generateJWTtest(userID string) string {
	jwtsecretkey := os.Getenv("JWTSECRETKEY")

	claim := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, _ := token.SignedString([]byte(jwtsecretkey))
	return signedToken
}

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode("test")

	tests := []struct {
		name   string
		token  string
		status int
	}{
		{"Valid token", generateJWTtest("1"), 200},
		// {"Invalid token", "", http.StatusBadRequest},
		// {"Invaild token", "bearrer" + "invalid.token.value", http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(AuthMiddleware())

			r.GET("/protect", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "OK"})
			})

			req, _ := http.NewRequest("GET", "/protect", nil)
			log.Println("これがtokenn:", tt.token)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)

		})
	}

}
