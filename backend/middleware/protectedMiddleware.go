package middleware

import (

	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"product/backend/util"


)
func Authentication(c *gin.Context){
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader,"Bearer "){
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
		c.Abort()

		return
	}
	tokenString := strings.TrimPrefix(authHeader,"Bearer ")

	claims,err := utils.VerifyToken(tokenString)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("claims",claims)
	c.Next()
}
