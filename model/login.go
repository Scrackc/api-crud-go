package model

import "github.com/golang-jwt/jwt/v4"

// Login .
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim .
type Claim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
