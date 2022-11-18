package middleware

import (
	"context"
	"learn-go/infrastructure"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FirebaseAuth(ctx *gin.Context) {
	firebaseAuth := infrastructure.AuthClient

	authorizationToken := ctx.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
		ctx.Abort()
		return
	}

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	ctx.Set("UUID", token.UID)
	ctx.Next()
}
