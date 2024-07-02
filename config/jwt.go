package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("hidsahdu19h912h82928982121h8")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
