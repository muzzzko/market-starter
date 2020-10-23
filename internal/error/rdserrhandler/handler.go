package rdserrhandler

import (
	"github.com/go-redis/redis"
	interror "market-starter/internal/error"
)


func NotFound(err error) bool {
	if err == nil {
		return false
	}

	if err == redis.Nil {
		return true
	}

	interror.Check(err)

	return false
}
