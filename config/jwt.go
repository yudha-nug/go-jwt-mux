package config

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("asadfafdasfaf343daf")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
