package controller

import (
	"errors"
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

// These secret keys will be environment variables in production
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
			"csrfToken": csrfToken,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not login"})
		return
	}

	// Set the jwt as HTTPOnly and Secure Cookie expires in 1 hour
	c.SetCookie("jwt", tokenString, 30, "", "", true, true)
	// Set custom header with csrfToken
	c.Header("CSRF-TOKEN", csrfToken)

	c.JSON(http.StatusOK, user)
}

// Logout is a gin HandlerFunc that handles requests to /api/logout
// logs out the signed in user by setting their jwt to be a negative expiration date which clears it from their cookies
func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -3600, "", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Dashboard is a gin HandlerFunc that handles requests to /api/dashboard
// LEVEL 1 requirement does not need to implement dashboard
// I am using this function to mimic a protected route for frontend testing.
// If JWT or csrfToken are invalid then it will return an error back to the frontend.
func GetDashboard(c *gin.Context) {
	csrfToken := c.Request.Header.Get("Csrf-Token")
	cookie, _ := c.Cookie("jwt")
	_, err := validateCsrfAndJWT(cookie, csrfToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// Everything valid, send some data back to the frontend.
	// Note: This should be some kind of dashboard data sent back but since Level 1
	// does not require dashboard I will just send a random object back
	c.JSON(http.StatusOK, struct{ Devices int }{Devices: 50})
}

// GetUser is a gin HandlerFunc that handles requests to /api/user
// It gets the currently logged in user depending on the csrfToken and jwt they sent to us
func GetUser(c *gin.Context) {
	csrfToken := c.Request.Header.Get("Csrf-Token")
	cookie, _ := c.Cookie("jwt")
	token, err := validateCsrfAndJWT(cookie, csrfToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// Find user
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	claims := token.Claims.(jwt.MapClaims)
	db.Where("id = ?", claims["iss"]).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// validateCsrfAndJWT is a helper function to help validate incoming requests that are not login or logout
// it will first check if the csrfToken is valid
// then it will check if the jwt is valid
// finally it will check if the jwt.csrfToken matches header.csrfToken
// returns a jwtToken if valid
func validateCsrfAndJWT(jwtCookie, csrfToken string) (*jwt.Token, error) {
	// Check if the csrfToken is valid
	if !xsrftoken.Valid(csrfToken, csrfTokenSecretKey, "", "") {
		return nil, errors.New("invalid csrfToken")
	}

	// Check if the jwt is valid
	token, err := jwt.Parse(jwtCookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid jwt")
	}

	// Check if the jwt's csrfToken matches the csrfToken from our header
	claims := token.Claims.(jwt.MapClaims)
	data := claims["data"].(map[string]interface{})
	csrfTokenFromClaims := data["csrfToken"].(string)
	if csrfTokenFromClaims != csrfToken {
		return nil, errors.New("jwt does not match csrf token")
	}

	return token, nil
}
