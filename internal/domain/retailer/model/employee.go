package model

type Employee struct {
	ID         int             `json:"id"`
	FirstName  string          `json:"firstName"`
	SecondName string          `json:"secondName"`
	Email      string          `json:"email"`
	Retailers  []*Retailer     `json:"retailers"`

	Password   string          `json:"-"`
	Hash       string		   `json:"-"`
}
