package middleware

import (
	"azarm-lending-backend/api"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

const JWT_SECRET = "There is no secret!"

func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(JWT_SECRET)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			api.NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			api.NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: "+err.Error())
			return
		}

		r.Header.Set("userId", claims.(jwt.MapClaims)["sub"].(string))
		next.ServeHTTP(w, r)
	})
}
