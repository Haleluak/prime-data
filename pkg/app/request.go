package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	prefix = "gin-go"
	UserIDKey        = prefix + "/user-id"
)

func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}

	return token
}

func GetUserID(c context.Context) string {
	userId := c.Value(UserIDKey)
	if userId == nil {
		return ""
	}
	return userId.(string)
}

// SetUserID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}