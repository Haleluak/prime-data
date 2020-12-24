package middleware

import (
	"github.com/gin-gonic/gin"
	"prime-data/pkg/app"
	gohttp "prime-data/pkg/http"
	"prime-data/pkg/http/wrapper"
	"prime-data/pkg/jwt"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	app.SetUserID(c, userID)
}

func UserAuthMiddleware(a jwt.IJWTAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := a.ParseUserID(app.GetToken(c), false)
		if err != nil {
			wrapper.Translate(c, gohttp.Response{Error: err})
			c.Abort()
			return
		}
		wrapUserAuthContext(c, userID)
		c.Next()
	}
}

func Authorize( sub, obj, act string)  gin.HandlerFunc {

}