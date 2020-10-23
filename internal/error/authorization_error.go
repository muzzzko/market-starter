package error

import "fmt"

type InvalidCredential struct {
	description string
	code string
}

func NewInvalidCredential() *InvalidCredential {
	return &InvalidCredential{
		description: "Invalid credential",
		code: "002-001",
	}
}

func (e *InvalidCredential) Error() string {
	return e.description
}

func (e *InvalidCredential) Code() string {
	return e.code
}

type InvalidSignedAlgorithm struct {
	description string
	code string
}

func NewInvalidSignedAlgorithm(algorithm interface{}) *InvalidSignedAlgorithm {
	return &InvalidSignedAlgorithm{
		description: fmt.Sprintf("Unexpected signing method: %v", algorithm),
		code: "002-002",
	}
}

func (e *InvalidSignedAlgorithm) Error() string {
	return e.description
}

func (e *InvalidSignedAlgorithm) Code() string {
	return e.code
}

type WrongEmailOrPassword struct {
	description string
	code string
}

func NewWrongEmailOrPassword() *WrongEmailOrPassword {
	return &WrongEmailOrPassword{
		description: "Wrong email or password",
		code: "002-003",
	}
}

func (e *WrongEmailOrPassword) Error() string {
	return e.description
}

func (e *WrongEmailOrPassword) Code() string {
	return e.code
}
