package dependency

import (
	"context"
	"market-starter/internal/domain/authorization/model"
	"time"
)

type SessionRepository interface {
	GetEmployeeSession(ctx context.Context, key string) *model.EmployeeSession
	SetEmployeeSession(ctx context.Context, key string, employeeSession *model.EmployeeSession, expiration time.Duration)
}