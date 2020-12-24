package wrapper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	gohttp "prime-data/pkg/http"

	"prime-data/pkg/errors"
)

const (
	DataField    = "data"
	TraceIDField = "trace_id"
	StatusField  = "status"
	CodeField    = "code"
	MessageField = "message"
)

type GinHandlerFn func(c *gin.Context) gohttp.Response

func Wrap(fn GinHandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		// handle req
		res := fn(c)

		Translate(c, res)
	}
}

func Translate(c *gin.Context, res gohttp.Response) {
	result := gin.H{}
	if _, ok := res.Error.(errors.CustomError); ok {
		status := int(errors.GetType(res.Error))
		result[StatusField] = status
		result[MessageField] = errors.GetMsg(status)
		result[CodeField] = errors.GetCode(status)
	}

	// get data
	if res.Data != nil {
		result[DataField] = res.Data
	}

	c.JSON(http.StatusOK, result)
}