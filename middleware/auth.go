package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"prime-data/enforce"
	"prime-data/pkg/app"
	"prime-data/pkg/errors"
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

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		m := c.Request.Method
		userId := app.GetUserID(c)

		ok, err := enforce.CasbinEnforce(e, userId, p, m)
		if err != nil {
			wrapper.Translate(c, gohttp.Response{Error: err})
			c.Abort()
			return
		}
		if !ok {
			wrapper.Translate(c, gohttp.Response{Error: errors.ErrMethodNotAllow})
			c.Abort()
			return
		}
		c.Next()
	}
}
