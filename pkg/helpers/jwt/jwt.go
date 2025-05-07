package jwt

import (
	"emp-app/app/dto"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Hard coded secret
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Generate access token and refresh token
func GenerateTokens(emp dto.EmployeeCreateRequest) (string, string, error) {
	// Access token valid for 15 minutes
	accessTokenClaims := &AuthCustomClaims{
		ID:    emp.ID,
		Email: emp.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessString, err := accessToken.SignedString(jwtSecret)

	log.Println("accesstoken : ", accessString)
	if err != nil {
		return "", "", fmt.Errorf("generation of access token failed : %w", err)
	}

	// Refresh token valid for 7 days
	refreshTokenClaims := &AuthCustomClaims{
		ID:    emp.ID,
		Email: emp.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshString, err := refreshToken.SignedString(jwtSecret)
	log.Println("refreshtoken: ", refreshString)
	if err != nil {
		return "", "", fmt.Errorf("generation of refresh token failed : %w", err)
	}

	return accessString, refreshString, nil

}
