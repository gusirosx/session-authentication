package models

import (
	"fmt"
	"os"
	"session-authentication/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var SecretKey []byte = []byte(os.Getenv("JWT_SECRET_KEY"))

// Login model function
func Login(cred entity.Credentials) (*entity.User, error) {

	//Get user function will be here
	var user *entity.User = &entity.UserDEMO

	if err := VerifyPassword(user.Password, cred.Password); err != nil {
		return &entity.User{}, err
	}
	if err := VerifyEmail(user.Email, cred.Email); err != nil {
		return &entity.User{}, err
	}

	// Expiration time of the token (kept it as 5 minutes)
	expirationTime := time.Now().Add(time.Minute * 5)
	// Create the JWT claims, which includes the username and expiry time
	claims := &entity.SignedDetails{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// JWT Creation
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return &entity.User{}, err
	}
	user.Token = tokenString

	return user, nil
}

func GetToken(user *entity.User) (string, error) {
	// Expiration time of the token (kept it as 1 hour)
	expirationTime := time.Now().Add(time.Hour * 1)
	// Create the JWT claims, which includes the username and expiry time
	claims := &entity.SignedDetails{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// JWT Creation
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Verify if both passwords are equal
func VerifyPassword(userPass, providedPass string) error {
	/* P.S: in production this passwords must be encrypted */
	if !strings.EqualFold(userPass, providedPass) {
		return fmt.Errorf("user password is incorrect")
	}
	return nil
}

// Checks if the email is already registered in the database
func VerifyEmail(email, providedEmail string) error {

	if !strings.EqualFold(email, providedEmail) {
		return fmt.Errorf("user e-mail is incorrect")
	}
	return nil
}

// Checks if the email is valid
func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email must be a non-empty string")
	}
	if parts := strings.Split(email, "@"); len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("malformed email string: %q", email)
	}
	return nil
}

/* This method will parse, validate, and return a error (if any) */
func ValidateToken(signedToken string) error {
	// Parse the JWT string and store the result in claims
	token, err := jwt.ParseWithClaims(signedToken, &entity.SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		return fmt.Errorf("invalid token signature")
	}
	claims, ok := token.Claims.(*entity.SignedDetails)
	if !ok || !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return fmt.Errorf("token is expired")
	}
	return nil
}

// This method will return an error if the token is invalid, or if the signature does not match
func GetTokenWithClaims(tokenValue string) (*jwt.Token, *entity.SignedDetails, error) {
	// Initialize a new instance of `Claims`
	claims := &entity.SignedDetails{}
	// Parse the JWT string and store the result in `claims`.
	token, err := jwt.ParseWithClaims(tokenValue, claims,
		func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
	if err != nil {
		return &jwt.Token{}, &entity.SignedDetails{}, err
	}
	return token, claims, nil
}
