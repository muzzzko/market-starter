package authorization

import (
	"context"
	"github.com/google/uuid"
	authmodel "market-starter/internal/domain/authorization/model"
	"market-starter/internal/domain/retailer/model"
)

type SessionManager interface {
	InitEmployeeSession(employee *model.Employee) *authmodel.EmployeeSession
	GetEmployeeSession(ctx context.Context, employeeID int, sessionID uuid.UUID) *authmodel.EmployeeSession
	SetEmployeeSession(ctx context.Context, employeeSession *authmodel.EmployeeSession)
}
