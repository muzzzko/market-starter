package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"market-starter/internal/domain/retailer/model"
	interror "market-starter/internal/error"
	"time"
)

func (s *HMAC) Generate(ctx context.Context, claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(s.secret)
	interror.Check(err)

	return signedString
}

func (s *HMAC) GetEmployeeClaims(ctx context.Context, employee *model.Employee) jwt.MapClaims {
	claims := jwt.MapClaims{}

	claims[emailClaim] = employee.Email
	claims[employeeIDClaim] = employee.ID

	now := time.Now()
	claims[expiredAtClaim] = now.Add(s.expiredAt).UTC().Unix()
	claims[issuedAtClaim] = now.UTC().Unix()

	return claims
}

func (s *HMAC) AddSessionIDClaims(claims jwt.MapClaims, sessionID uuid.UUID) {
	claims[sessionIDClaim] = sessionID.String()
}
