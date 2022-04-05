package entity

import "github.com/golang-jwt/jwt"

// User Credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Claims
type SignedDetails struct {
	ID        string `json:"UID"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	jwt.StandardClaims
}
