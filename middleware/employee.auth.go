package middleware

import (
	jwtpackage "emp-app/pkg/helpers/jwt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := jwtpackage.ExtractTokenFromHeader(r)

		// Checking token is empty or not
		if tokenString == "" {
			http.Error(w, " Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// Checking token blacklisted or not
		if jwtpackage.IsTokenBlackListed(tokenString) {
			http.Error(w, "user need to login..(blacklisted token)", http.StatusUnauthorized)
			return
		}


		token, err := jwt.ParseWithClaims(tokenString, &jwtpackage.AuthCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return jwtpackage.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token or Expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*jwtpackage.AuthCustomClaims)
		if !ok || claims.Role != "admin" {
			http.Error(w, "Forbidden: admin access required", http.StatusForbidden)
			return
		}

		// If everything went ok, call next handler
		next.ServeHTTP(w, r)

	})
}
