// Package tokens contains tools for token-based authentication
package tokens

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/moooll/dogs-service/internal/models"
)

// JWTClaims represents jwt claims
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SigningKey is the secret to sign the tokens
type SigningKey struct {
	secret []byte
}

// NewSigningKey creates new signing key
func NewSigningKey(secret []byte) *SigningKey {
	return &SigningKey{
		secret: secret,
	}
}

// GenerateToken generates new JWT-token and signs it, returning token string and an error
func (s *SigningKey) GenerateToken(u models.User) (string, error) {
	claims := JWTClaims{
		u.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(s.secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

// ValidateToken checks if the token string is valid and return username and error
func (s *SigningKey) ValidateToken(tokenStr string) (string, error) {
	var claims JWTClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, errors.New("unexpected signing method: " + fmt.Sprint(token.Header["alg"]))
		}
		return s.secret, nil
	})
	switch {
	case err != nil:
		return "", err
	case !token.Valid:
		return "", errors.New("invalid token")
	}

	return claims.Username, nil
}
