package dependency

import (
	"context"
	"market-starter/internal/domain/retailer/model"
)

type RetailerEmployeeRepository interface {
	InsertEmployee(ctx context.Context, employee *model.Employee)
	GetEmployee(ctx context.Context, employeeID int) *model.Employee
	GetEmployeeByEmail(ctx context.Context, email string) *model.Employee
	InsertRetailerForRetailerEmployee(ctx context.Context, employee *model.Employee, retailer *model.Retailer)
}

type RetailerEmployeeRoleRepository interface {
	GetRoleByName(ctx context.Context, roleName string) *model.RetailerEmployeeRole
}