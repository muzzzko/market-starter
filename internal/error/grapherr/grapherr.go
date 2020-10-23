package grapherr

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	interror "market-starter/internal/error"
)

func Serialize(err interror.Error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": err.Code(),
		},
	}
}
