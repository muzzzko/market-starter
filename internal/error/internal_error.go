package error

type InternalError struct {
	description string
	code string
}

func NewInternalError() *InternalError {
	return &InternalError{
		description: "Something went wrong. Try again",
		code: "000-001",
	}
}

func (e *InternalError) Error() string {
	return e.description
}

func (e *InternalError) Code() string {
	return e.code
}