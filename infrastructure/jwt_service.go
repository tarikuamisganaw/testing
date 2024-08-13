package infrastructure

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTService defines the interface for generating and validating JWTs.
type JWTService interface {
	GenerateJWT(username, role string) (string, error)
	ValidateJWT(tokenString string) (*jwt.Token, error)
}

// jwtService is the implementation of JWTService.
type jwtService struct {
	secretKey []byte
}

// NewJWTService creates a new instance of jwtService with the provided secret key.
func NewJWTService(secretKey string) JWTService {
	return &jwtService{secretKey: []byte(secretKey)}
}

// GenerateJWT generates a JWT token for a given username and role.
func (j *jwtService) GenerateJWT(username, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// ValidateJWT validates a JWT token string and returns the parsed token.
func (j *jwtService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
