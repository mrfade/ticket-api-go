package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrfade/ticket-api-go/helpers"
	"github.com/mrfade/ticket-api-go/initializers"
	"github.com/mrfade/ticket-api-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get the email and password from the req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		helpers.ErrorJSON(c, http.StatusBadRequest, "Request body is invalid")
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		helpers.ErrorJSON(c, http.StatusInternalServerError, "Failed to hash the password")
		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		helpers.ErrorJSON(c, http.StatusInternalServerError, "Failed to create the user")
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// get the email and pass from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		helpers.ErrorJSON(c, http.StatusBadRequest, "Request body is invalid")
		return
	}

	// look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		helpers.ErrorJSON(c, http.StatusNotFound, "User not found")
		return
	}

	// compare the hash and given password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		helpers.ErrorJSON(c, http.StatusBadRequest, "Password is invalid")
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		helpers.ErrorJSON(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
