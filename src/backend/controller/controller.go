package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/teeaaspoon/react-go-login/src/backend/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/xsrftoken"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

const jwtSecretKey = "jwtSecretKey"
const csrfTokenSecretKey = "csrfTokenSecretKey"

// Login is a gin HandlerFunc that handles requests to /api/login
// It verifies our input and tries to find the user with matching email in our database
// Then it tries to compare the hashes of the passwords.
// If all is successful it will generate a csrfToken and sign a jwt and send it back as the response
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var data map[string]string

	// Validate our input
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	var user models.User
	db.Where("email = ?", data["email"]).First(&user)
	fmt.Println(user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password"})
		return
	}

	// Generate csrfToken
	csrfToken := xsrftoken.Generate(csrfTokenSecretKey, "", "")

	// Issue and sign jwt
	claims := &jwt.MapClaims{
		"iss": strconv.Itoa(int(user.Id)),
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]string{
			"xsrfToken": csrfToken,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not login"})
		return
	}

	// Set the jwt as HTTPOnly and Secure Cookie
	c.SetCookie("jwt", tokenString, 3600, "", "", false, true)
	// Set custom header with csrfToken
	c.Header("CSRF-TOKEN", csrfToken)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Logout is a gin HandlerFunc that handles requests to /api/logout
// logs out the signed in user by setting their jwt to be a negative expiration date which clears it from their cookies
func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
