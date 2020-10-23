package model

type Retailer struct {
	ID             int                         `json:"id"`
	Name           string                      `json:"name"`
	EmployeeRole   *RetailerEmployeeRole        `json:"EmployeeRole"`

	Secret 		   string 					   `json:"-"`
	IsNew          bool 					   `json:"-"`
}


type NewRetailerEmployee struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type RetailerEmployeeRole struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
