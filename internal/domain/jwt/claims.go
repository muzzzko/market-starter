package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"market-starter/internal/logger"
)

const (
	emailClaim = "email"
	employeeIDClaim = "employee_id"

	sessionIDClaim = "session_id"

	expiredAtClaim = "exp"
	issuedAtClaim = "iat"
)

func GetEmployeeIDFromClaims(ctx context.Context, claims jwt.MapClaims) (int, bool) {
	employeeID, ok := claims["employee_id"].(float64)
	if !ok {
		logger.WithContext(ctx).Info(logger.EmployeeIDClaimIsMissed)
		return 0, ok
	}

	return int(employeeID), ok
}

func GetSessionIDFromClaims(ctx context.Context, claims jwt.MapClaims) (uuid.UUID, bool) {
	sessionIDString, ok := claims["session_id"].(string)
	if !ok {
		logger.WithContext(ctx).Info(logger.SessionIDClaimIsMissed)
		return uuid.UUID{}, ok
	}
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		logger.WithContext(ctx).With(zap.Error(err)).Info(logger.InvalidSessionID)
		return uuid.UUID{}, false
	}

	return sessionID, true
}