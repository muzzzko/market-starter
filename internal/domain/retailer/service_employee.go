package retailer

import (
	"context"
	"crypto/rand"
	"fmt"
	"go.uber.org/zap"
	"market-starter/internal/domain/retailer/dependency"
	"market-starter/internal/domain/retailer/model"
	interror "market-starter/internal/error"
	"market-starter/internal/logger"
)

type EmployeeServiceImp struct {
	employeeRepository dependency.RetailerEmployeeRepository
	employeeRoleRepository dependency.RetailerEmployeeRoleRepository

	passwordService dependency.PasswordService
}

func NewEmployeeService(
	employeeRepository dependency.RetailerEmployeeRepository,
	employeeRoleRepository dependency.RetailerEmployeeRoleRepository,
	passwordService dependency.PasswordService,
) *EmployeeServiceImp {

	return &EmployeeServiceImp {
		employeeRepository: employeeRepository,
		employeeRoleRepository: employeeRoleRepository,
		passwordService: passwordService,
	}
}

func (s *EmployeeServiceImp) CreateEmployee(ctx context.Context, employee *model.Employee) interror.Error {
	if employee := s.employeeRepository.GetEmployeeByEmail(ctx, employee.Email); employee != nil {
		return interror.NewEmployeeEmailAlreadyUsed()
	}

	employee.Hash = s.passwordService.GetHash(employee.Password)

	s.employeeRepository.InsertEmployee(ctx, employee)

	return nil
}

func (s *EmployeeServiceImp) GetEmployee(ctx context.Context, employeeID int) *model.Employee {
	return s.employeeRepository.GetEmployee(ctx, employeeID)
}

func (s *EmployeeServiceImp) CheckPasswordAndReturnEmployeeByEmail(ctx context.Context, email string, password string) (*model.Employee, interror.Error) {
	employee := s.employeeRepository.GetEmployeeByEmail(ctx, email)
	if employee == nil {
		logger.WithContext(ctx).With(
			zap.String(logger.EmailField, email),
		).Info(logger.EmployeeNotFound)

		return nil, interror.NewWrongEmailOrPassword()
	}

	if s.passwordService.HashIsEqualToPassword(employee.Password, password) {
		logger.WithContext(ctx).With(
			zap.String(logger.EmailField, email),
		).Info(logger.WrongPassword)

		return nil, interror.NewWrongEmailOrPassword()
	}

	return employee, nil
}

func (s *EmployeeServiceImp) AddRetailerForRetailerEmployee(
	ctx context.Context,
	employee *model.Employee,
	retailer *model.Retailer,
	roleName string,
) interror.Error {

	role := s.employeeRoleRepository.GetRoleByName(ctx, roleName)
	if role == nil {
		return interror.NewEmployeeRoleNotFound()
	}

	retailer.EmployeeRole = role
	retailer.Secret = generateRetailerSecret()

	s.employeeRepository.InsertRetailerForRetailerEmployee(ctx, employee, retailer)

	employee.Retailers = append(employee.Retailers, retailer)

	return nil
}

func generateRetailerSecret() string {
	secretLength := 32
	secret := make([]byte, secretLength)

	_, err := rand.Read(secret)
	interror.Check(err)

	return fmt.Sprintf("%x", secret)
}