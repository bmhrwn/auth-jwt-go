package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jeypc/go-jwt-mux/config"
	"github.com/jeypc/go-jwt-mux/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}

		// Mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaim{}
		// Parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// Token Invalid
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				// Token Expired
				response := map[string]string{"message": "Unauthorized, Token Expired!"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
