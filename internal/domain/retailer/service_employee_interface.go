package retailer

import (
	"context"
	"market-starter/internal/domain/retailer/model"
	interror "market-starter/internal/error"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee *model.Employee) interror.Error
	GetEmployee(ctx context.Context, employeeID int) *model.Employee
	CheckPasswordAndReturnEmployeeByEmail(ctx context.Context, email string, password string) (*model.Employee, interror.Error)
	AddRetailerForRetailerEmployee(
		ctx context.Context,
		employee *model.Employee,
		retailer *model.Retailer,
		roleName string,
	) interror.Error
}
