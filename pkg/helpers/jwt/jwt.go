package jwt

import (
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

// JWT secret string
var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Generate access token and refresh token
func GenerateTokens(email string) (string, string, error) {
	// Access token valid for 15 minutes
	accessTokenClaims := &AuthCustomClaims{

		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessString, err := accessToken.SignedString(JwtSecret)

	log.Println("accesstoken : ", accessString)
	if err != nil {
		return "", "", fmt.Errorf("generation of access token failed : %w", err)
	}

	// Refresh token valid for 7 days
	refreshTokenClaims := &AuthCustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshString, err := refreshToken.SignedString(JwtSecret)
	log.Println("refreshtoken: ", refreshString)
	if err != nil {
		return "", "", fmt.Errorf("generation of refresh token failed : %w", err)
	}

	return accessString, refreshString, nil

}

// For blacklisting token
func BlackListToken(token string) error {
	mu.Lock()
	defer mu.Unlock()
	BlackListedTokens[token] = true
	return nil
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
