package main

import (
	"crypto/rand"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateTokens(userID string) TokensPair {
	refreshToken, refreshTokenHash := GenerateRFTokenWithHash()
	accessToken := GenerateAccessToken(refreshTokenHash, userID)

	if !IsUserExist(userID) {
		AddUser(userID, refreshTokenHash)
	} else {
		UpdateUserRFTokenHash(userID, refreshTokenHash)
	}

	return TokensPair{AccessToken: accessToken, RefreshToken: refreshToken}
}

func GenerateRFTokenWithHash() ([]byte, []byte) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}

	refreshTokenHash, err := bcrypt.GenerateFromPassword(randomBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return randomBytes, refreshTokenHash
}

func GenerateAccessToken(parentHash []byte, userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &TokenClaims{
		RFTokenHash: parentHash,
		UserID:      userID,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}
