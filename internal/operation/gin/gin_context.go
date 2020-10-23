package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	interror "market-starter/internal/error"
)

const (
	ginContextKey = "GinContextKey"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, interror.Error) {
	ginContext := ctx.Value(ginContextKey)
	if ginContext == nil {
		//todo: add logger
		//err := fmt.Errorf("could not retrieve gin.Context")
		return nil, interror.NewInternalError()
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		//todo: add logger
		//err := fmt.Errorf("gin.Context has wrong type")
		return nil, interror.NewInternalError()
	}
	return gc, nil
}
