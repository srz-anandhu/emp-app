package jwt

import (
	"emp-app/app/dto"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

// To blacklist tokens
var BlackListedTokens = make(map[string]bool)

// For safe read/writing
var mu sync.RWMutex

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

// For blacklisting token
func BlackListToken(token string) {
	mu.Lock()
	defer mu.Unlock()
	BlackListedTokens[token] = true
}

// Check token blacklisted or not
func IsTokenBlackListed(token string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return BlackListedTokens[token]
}

// Get Token from authorization header
func ExtractTokenFromHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}

	return token
}
