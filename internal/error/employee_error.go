package error

type EmployeeEmailAlreadyUsed struct {
	description string
	code string
}

func NewEmployeeEmailAlreadyUsed() *EmployeeEmailAlreadyUsed {
	return &EmployeeEmailAlreadyUsed {
		description: "Email already used",
		code: "001-001",
	}
}

func (e *EmployeeEmailAlreadyUsed) Error() string {
	return e.description
}

func (e *EmployeeEmailAlreadyUsed) Code() string {
	return e.code
}

type EmployeeRoleNotFound struct {
	description string
	code string
}

func NewEmployeeRoleNotFound() *EmployeeRoleNotFound {
	return &EmployeeRoleNotFound {
		description: "Employee role not found",
		code: "001-002",
	}
}

func (e *EmployeeRoleNotFound) Error() string {
	return e.description
}

func (e *EmployeeRoleNotFound) Code() string {
	return e.code
}
