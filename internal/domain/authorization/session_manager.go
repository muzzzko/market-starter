package authorization

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"market-starter/internal/domain/authorization/dependency"
	authmodel "market-starter/internal/domain/authorization/model"
	"market-starter/internal/domain/retailer/model"
	"market-starter/internal/logger"
	"time"
)

const (
	employeeSessionKey = "employee_session:%d:%s"
)



type SessionManagerImp struct {
	sessionRepository dependency.SessionRepository
	employeeSessionExpiration time.Duration
}



var (
	sessionManager *SessionManagerImp
)

func NewSessionManagerImp(sessionRepository dependency.SessionRepository, employeeSessionExpiration time.Duration) *SessionManagerImp {
	if sessionManager == nil {
		sessionManager = &SessionManagerImp{
			sessionRepository:         sessionRepository,
			employeeSessionExpiration: employeeSessionExpiration,
		}
	}

	return sessionManager
}



func (m *SessionManagerImp) InitEmployeeSession(employee *model.Employee) *authmodel.EmployeeSession {
	employeeSession := &authmodel.EmployeeSession{
		EmployeeID: employee.ID,
		SessionID: uuid.New(),
	}

	return employeeSession
}

func (m *SessionManagerImp) GetEmployeeSession(ctx context.Context, employeeID int, sessionID uuid.UUID) *authmodel.EmployeeSession {
	key := fmt.Sprintf(employeeSessionKey, employeeID, sessionID)

	employeeSession := m.sessionRepository.GetEmployeeSession(ctx, key)
	if employeeSession == nil {
		logger.WithContext(ctx).With(
			zap.Int(logger.EmployeeIDField, int(employeeID)),
			zap.String(logger.SessionIDField, sessionID.String()),
		).Info(logger.SessionNotFound)

		return nil
	}

	employeeSession.EmployeeID = employeeID
	employeeSession.SessionID = sessionID

	return employeeSession
}

func (m *SessionManagerImp) SetEmployeeSession(ctx context.Context, employeeSession *authmodel.EmployeeSession) {
	key := fmt.Sprintf(employeeSessionKey, employeeSession.EmployeeID, employeeSession.SessionID)

	m.sessionRepository.SetEmployeeSession(ctx, key, employeeSession, m.employeeSessionExpiration)
}