package graph

import (
	"market-starter/internal/domain/authorization"
	"market-starter/internal/domain/jwt"
	"market-starter/internal/domain/retailer"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	employeeService retailer.EmployeeService
	jwtService jwt.Service
	sessionManager authorization.SessionManager
}

func NewResolver(
	employeeService retailer.EmployeeService,
	jwtService jwt.Service,
	sessionManager authorization.SessionManager,
) *Resolver {

	return &Resolver{
		employeeService: employeeService,
		jwtService: jwtService,
		sessionManager: sessionManager,
	}
}
