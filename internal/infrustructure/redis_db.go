package infrustructure

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
	"market-starter/config"
	interror "market-starter/internal/error"
	"time"
)

var (
	rdb *redis.Client
)

func NewRedisDB(cfg *config.Config) *redis.Client {
	if rdb == nil {
		options := &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			Password:     cfg.RedisPassword,
			DialTimeout:  time.Duration(cfg.RedisTimeout) * time.Millisecond,
			ReadTimeout:  time.Duration(cfg.RedisReadTimeout) * time.Millisecond,
			WriteTimeout: time.Duration(cfg.RedisWriteTimeout) * time.Millisecond,
		}

		rdb = redis.NewClient(options)

		err := rdb.Ping().Err()
		interror.Check(err)

		rdb.AddHook(nrredis.NewHook(options))
	}

	return rdb
}