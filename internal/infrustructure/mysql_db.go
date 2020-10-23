package infrustructure

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"market-starter/config"
	interror "market-starter/internal/error"
	"time"
)

var (
	db *sql.DB
)

func NewMySQLDB(cfg *config.Config) *sql.DB {
	if db == nil {
		mysqlCfg := &mysql.Config{
			Addr:                 fmt.Sprintf("%s:%s", cfg.DatabaseHost, cfg.DatabasePort),
			User:                 cfg.DatabaseUser,
			Passwd:               cfg.DatabasePassword,
			Timeout:              time.Duration(cfg.DatabaseTimeout) * time.Millisecond,
			WriteTimeout:         time.Duration(cfg.DatabaseWriteTimeout) * time.Millisecond,
			ReadTimeout:          time.Duration(cfg.DatabaseReadTimeout) * time.Millisecond,
			AllowNativePasswords: true,
		}

		var err error
		db, err = sql.Open("nrmysql", mysqlCfg.FormatDSN())
		interror.Check(err)

		err = db.Ping()
		interror.Check(err)
	}

	return db
}
