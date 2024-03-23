package main

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID          string `json:"id"`
	RFTokenHash string `json:"rf_token_hash" bson:"rf_token_hash"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	RFTokenHash []byte `json:"rf_token_hash"`
	UserID      string `json:"userID"`
}

type TokensPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken []byte `json:"refresh_token"`
}
