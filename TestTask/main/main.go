package main

import "github.com/gin-gonic/gin"

const secretKey string = "secret-key-from-env-var"

var client, cancel = ConnectDB()

func main() {
	defer cancel()
	router := gin.Default()

	router.GET("/get_tokens", GetTokens)
	router.POST("/refresh", RefreshTokens)
		

	router.Run("localhost:8000")
}
