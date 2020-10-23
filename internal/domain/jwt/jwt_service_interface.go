package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"market-starter/internal/domain/retailer/model"
)

type Service interface {
	GetEmployeeClaims(ctx context.Context, employee *model.Employee) jwt.MapClaims
	Generate(ctx context.Context, claims jwt.MapClaims) string
	AddSessionIDClaims(claims jwt.MapClaims, sessionID uuid.UUID)

	ValidateAndReturnClaims(ctx context.Context, signedString string) (jwt.MapClaims, error)
}
