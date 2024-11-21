package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gocourse.com/restapi/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization") // getting the Authorization header from the request

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is required"})
		return
	}

	userId, err := utils.VerifyToken(token) // verifying the token

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized: Invalid token"})
		return
	}

	context.Set("userId", userId) // setting the userId in the context for later use

	context.Next()
}
