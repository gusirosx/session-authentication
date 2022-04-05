package entity

// User Struct
type User struct {
	ID        string `json:"userID"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Token     string `json:"token"`
}

// Example User
var UserDEMO User = User{
	ID:        "0001",
	UserName:  "gusirox",
	FirstName: "Gustavo",
	LastName:  "Rodrigues",
	Password:  "123456",
	Email:     "gusirosx@email.com",
	Phone:     "+5534900000001",
}
