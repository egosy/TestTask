package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetTokens(c *gin.Context) {
	userID := c.Query("userID")
	c.IndentedJSON(200, GenerateTokens(userID))
}

func RefreshTokens(c *gin.Context) {
	var tokensPair TokensPair
	if err := c.BindJSON(&tokensPair); err != nil {
		c.JSON(401, gin.H{"error": "invalid JSON or token structure"})
		return
	}

	tokenClaims := &TokenClaims{}
	_, err := jwt.ParseWithClaims(tokensPair.AccessToken, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"error": "invalid access token"})
		return
	} else if bcrypt.CompareHashAndPassword(tokenClaims.RFTokenHash, tokensPair.RefreshToken) != nil {
		c.JSON(401, gin.H{"error": "invalid refresh token"})
		return
	} else if !IsHashExistsInDB(tokenClaims.RFTokenHash) {
		c.JSON(401, gin.H{"error": "no such refresh token in db"})
		return
	}

	c.IndentedJSON(200, GenerateTokens(tokenClaims.UserID))
}
