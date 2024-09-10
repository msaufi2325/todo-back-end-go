package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Issuer      string
	Audience    string
	Secret      string
	TokenExpiry time.Duration
	RefreshExpiry time.Duration
	CookieDomain string
	CookiePath string
	CookieName string
}

type jwtUser struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

type TokenPairs struct {
	Token string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}
