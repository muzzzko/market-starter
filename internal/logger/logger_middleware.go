package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddLoggerToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIP := c.Request.RemoteAddr
		referer := c.Request.Referer()
		RequestID := c.Request.Header.Get("Request-ID")

		ctx := NewContext(c.Request.Context(),
			zap.String(UserIPField, userIP),
			zap.String(RefererField, referer),
			zap.String(RequestIDField, RequestID),
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}



