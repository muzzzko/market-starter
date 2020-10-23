package authorization

import (
	"context"
	"github.com/gin-gonic/gin"
)

func AddTokenToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo add var for cookie name
		cookie, err := c.Request.Cookie("auth")

		// Allow unauthenticated users in
		if err != nil || cookie == nil {
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), tokenContextKey, cookie.Value)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}


