package tollbooth_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/unix-world/go-modules/didip/tollbooth"
	"github.com/unix-world/go-modules/didip/tollbooth/config"
)

func LimitHandler(limiter *config.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(limiter, c.Request)
		if httpError != nil {
			c.String(httpError.StatusCode, httpError.Message)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
