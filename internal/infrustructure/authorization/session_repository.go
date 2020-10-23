package authorization

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"market-starter/internal/domain/authorization/model"
	interror "market-starter/internal/error"
	"market-starter/internal/error/rdserrhandler"
	"time"
)

type RedisEmployeeSessionRepository struct {
	rdb *redis.Client

}



var (
	redisEmployeeSessionRepository *RedisEmployeeSessionRepository
)

func NewRedisEmployeeSessionRepository(rdb *redis.Client) *RedisEmployeeSessionRepository {
	if redisEmployeeSessionRepository == nil {
		redisEmployeeSessionRepository = &RedisEmployeeSessionRepository{
			rdb: rdb,
		}
	}

	return redisEmployeeSessionRepository
}



func (r *RedisEmployeeSessionRepository) GetEmployeeSession(ctx context.Context, key string) *model.EmployeeSession {
	result, err := r.rdb.WithContext(ctx).Get(key).Result()
	if rdserrhandler.NotFound(err) {
		return nil
	}

	employeeSession := &model.EmployeeSession{}
	err = json.Unmarshal([]byte(result), employeeSession)
	interror.Check(err)

	return employeeSession
}

func (r *RedisEmployeeSessionRepository) SetEmployeeSession(
	ctx context.Context,
	key string,
	employeeSession *model.EmployeeSession,
	expiration time.Duration,
) {
	data, err := json.Marshal(employeeSession)
	interror.Check(err)

	err = r.rdb.WithContext(ctx).Set(key, data, expiration).Err()
	interror.Check(err)
}