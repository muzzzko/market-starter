package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	interror "market-starter/internal/error"
	"market-starter/internal/logger"
)

func (s *HMAC) ValidateAndReturnClaims(ctx context.Context, signedString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, interror.NewInvalidSignedAlgorithm(token.Header["alg"])
		}

		return s.secret, nil
	})

	if err != nil {
		logger.WithContext(ctx).Info(logger.TokenInvalid)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	logger.WithContext(ctx).Info(logger.TokenInvalid)
	return nil, err
}
