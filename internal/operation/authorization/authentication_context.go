package authorization

import (
	"context"
	"market-starter/internal/domain/authorization/model"
)

const (
	employeeSessionContextKey = "employeeSessionContextKey"
	tokenContextKey = "tokenContextKey"
)

func EmployeeSessionFromContext(ctx context.Context) *model.EmployeeSession {
	employeeSession, _ := ctx.Value(employeeSessionContextKey).(*model.EmployeeSession)
	//todo: handle error

	return employeeSession
}

func AddEmpoyeeSessionToContext(ctx context.Context, employeeSession *model.EmployeeSession) context.Context {
	return context.WithValue(ctx, employeeSessionContextKey, employeeSession)
}