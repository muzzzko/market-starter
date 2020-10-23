package newrelic

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		txn := nrgin.Transaction(c)
		if txn != nil {
			ctx := newrelic.NewContext(c.Request.Context(), txn)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}