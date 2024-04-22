package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const USERNAME string = "admin"
const PASSWORD string = "admin"

// All requests will go through this function
func AuthHandler(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	// Checks if request has basic auth
	if !ok {
		err := errors.New("authentication failed")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Checks if username and password are correct
	if username != USERNAME || password != PASSWORD {
		err := errors.New("incorrect username or password")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
}
