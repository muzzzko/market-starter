package config

import (
	"github.com/jessevdk/go-flags"
	imperror "market-starter/internal/error"
)

type Config struct {
	DatabaseHost string `long:"database-host" env:"DATABASE_HOST" required:"true"`
	DatabasePort string `long:"database-post" env:"DATABASE_PORT" required:"true"`
	DatabaseUser string `long:"database-user" env:"DATABASE_USER" required:"true"`
	DatabasePassword string `long:"database-password" env:"DATABASE_PASSWORD" required:"true"`
	DatabaseTimeout int `long:"database-timeout" env:"DATABASE_TIMEOUT" `
	DatabaseWriteTimeout int `long:"database-write-timeout" env:"DATABASE_WRITE_TIMEOUT" default:"1000" description:"I/O write timeout (ms)"`
	DatabaseReadTimeout int `long:"database-read-timeout" env:"DATABASE_READ_TIMEOUT" default:"1000" description:"I/O read timeout (ms)"`
	
	RedisHost string `long:"redis-host" env:"REDIS_HOST" required:"true"`
	RedisPort string `long:"redis-post" env:"REDIS_PORT" required:"true"`
	RedisPassword string `long:"redis-password" env:"REDIS_PASSWORD" required:"true"`
	RedisTimeout int `long:"redis-timeout" env:"REDIS_TIMEOUT" `
	RedisWriteTimeout int `long:"redis-write-timeout" env:"REDIS_WRITE_TIMEOUT" default:"1000" description:"I/O write timeout (ms)"`
	RedisReadTimeout int `long:"redis-read-timeout" env:"REDIS_READ_TIMEOUT" default:"1000" description:"I/O read timeout (ms)"`

	AuthorizationSecret string `long:"auth-secret" env:"AUTH_SECRET" required:"true" description:"auth secret for signing employee token"`
	SessionExpiration int `long:"session-exp" env:"SESSION_EXPIRATION" default:"86000" description:"session ttl (s)"`

	LoggerProgramNameField string `long:"program-name" env:"PROGRAM_NAME" default:"market-starter"`
	LoggerPathToLogFile string `long:"log-file" env:"PATH_TO_LOG_FILE" default:"/tmp/log/fifo.json"`

	NewRelicAppName string `long:"newrelic-app-name" env:"NEWRELIC_APP_NAME" default:"Market starter"`
	NewRelicLicense string `long:"newrelic-license" env:"NEWRELIC_LICENSE"`
}

var (
	cfg *Config
)

func NewConfig() *Config {
	if cfg == nil {
		cfg = &Config{}

		_, err := flags.ParseArgs(cfg, nil)
		imperror.Check(err)
	}

	return cfg
}
