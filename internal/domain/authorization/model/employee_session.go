package model

import "github.com/google/uuid"

type EmployeeSession struct {
	EmployeeID int `json:"-"`
	SessionID uuid.UUID `json:"-"`
}


