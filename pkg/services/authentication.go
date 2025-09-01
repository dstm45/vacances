// Package services contient toute la logique qu serra utilisée dans les controllers
package services

import (
	"github.com/dstm45/vacances/pkg/api"
	"github.com/dstm45/vacances/pkg/database"
	"github.com/golang-jwt/jwt/v5"
)

// AuthClaim defines the structure of the JWT claims.
type AuthService struct {
	db *database.Queries
}

// NewAuthService creates a new AuthService.
func NewAuthService(db *database.Queries, config api.Config) *AuthService {
	return &AuthService{
		db: db,
	}
}

// CreateAccessToken generates a new JWT access token for a user.
func (s *AuthService) CreateAccessToken(user database.User, config api.Config) (string, error) {
	claim := jwt.RegisteredClaims{
		Subject: user.UUID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(config.Secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// CreateRefreshToken generates a new refresh token for a user.
func (s *AuthService) CreateRefreshToken(user database.User) {
}

// ValidateAccessToken parses and validates a JWT access token string.
// It returns the custom claims if the token is valid.
func (s *AuthService) ValidateAccessToken(tokenString string) {
}

// RefreshAccessToken validates a refresh token and issues a new access token.
func (s *AuthService) RefreshAccessToken(refreshTokenString string) {
}

// BlacklistToken adds a token's ID to a blacklist (e.g., in Redis or a database)
// to prevent its reuse after logout.
func (s *AuthService) BlacklistToken(tokenID string) {
}

// TODO: Modifier les signatures des fonction et implémenter ces fonctions.
