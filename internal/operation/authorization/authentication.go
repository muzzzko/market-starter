package authorization

import (
	"context"
	"market-starter/internal/domain/authorization"
	"market-starter/internal/domain/authorization/model"
	"market-starter/internal/domain/jwt"
	"market-starter/internal/logger"
)

type Authenticator interface {
	AuthenticateAndReturnEmployeeSession(ctx context.Context) *model.EmployeeSession
}

type AuthenticatorImp struct {
	jwtService jwt.Service
	sessionManager authorization.SessionManager
}



func NewAuthenticator(jwtService jwt.Service, sessionManager authorization.SessionManager) *AuthenticatorImp {
	return &AuthenticatorImp{
		jwtService: jwtService,
		sessionManager: sessionManager,
	}
}



func(a *AuthenticatorImp) AuthenticateAndReturnEmployeeSession(ctx context.Context) *model.EmployeeSession {
	signedString := getTokenFromContext(ctx)
	claims, err := a.jwtService.ValidateAndReturnClaims(ctx, signedString)
	if err != nil {
		return nil
	}

	employeeID, ok := jwt.GetEmployeeIDFromClaims(ctx, claims)
	if !ok {
		return nil
	}

	sessionID, ok := jwt.GetSessionIDFromClaims(ctx, claims)
	if !ok {
		return nil
	}

	employeeSession := a.sessionManager.GetEmployeeSession(ctx, int(employeeID), sessionID)
	if employeeSession == nil {
		return nil
	}

	return employeeSession
}

func getTokenFromContext(ctx context.Context) string {
	signedString, ok := ctx.Value(tokenContextKey).(string)
	if !ok {
		logger.WithContext(ctx).Info(logger.TokenMissed)
		return ""
	}

	return signedString
}