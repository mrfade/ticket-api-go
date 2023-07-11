package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrfade/ticket-api-go/initializers"
	"github.com/mrfade/ticket-api-go/models"
)

func RequireAuth(c *gin.Context) {
	// get the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode and validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// check the exp
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// find the user with token sub
	var user models.User
	initializers.DB.Preload("Roles").First(&user, claims["sub"])

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// attach to req
	c.Set("user", user)

	// continue
	c.Next()
}
