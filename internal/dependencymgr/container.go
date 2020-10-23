package dependencymgr

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"market-starter/config"
	"market-starter/internal/domain/authorization"
	"market-starter/internal/domain/jwt"
	"market-starter/internal/domain/password"
	retailersrv "market-starter/internal/domain/retailer"
	"market-starter/internal/graphql/graph"
	"market-starter/internal/infrustructure"
	infauth "market-starter/internal/infrustructure/authorization"
	interror "market-starter/internal/error"
	"market-starter/internal/infrustructure/retailer"
	"market-starter/internal/logger"
	operauth "market-starter/internal/operation/authorization"
	"os"
	"time"
)

// todo: replace on dig when dig.As feature will be released

func NewResolver() *graph.Resolver {
	cfg := config.NewConfig()

	db := infrustructure.NewMySQLDB(cfg)

	rdb := infrustructure.NewRedisDB(cfg)

	passwordService := password.NewSHA256()
	jwtService := jwt.NewHMAC([]byte(cfg.AuthorizationSecret), time.Duration(cfg.SessionExpiration) * time.Second)

	sessionEmployeeRepository := infauth.NewRedisEmployeeSessionRepository(rdb)
	sessionManager := authorization.NewSessionManagerImp(sessionEmployeeRepository, time.Duration(cfg.SessionExpiration) * time.Second)

	employeeRepository := retailer.NewMySQLEmployeeRepository(db)
	employeeRoleRepository := retailer.NewEmployeeRepository(db)
	employeeService := retailersrv.NewEmployeeService(employeeRepository, employeeRoleRepository, passwordService)

	resolver := graph.NewResolver(employeeService, jwtService, sessionManager)

	return resolver
}

func NewAuthenticator() operauth.Authenticator {
	cfg := config.NewConfig()

	rdb := infrustructure.NewRedisDB(cfg)

	jwtService := jwt.NewHMAC([]byte(cfg.AuthorizationSecret), time.Duration(cfg.SessionExpiration) * time.Second)

	sessionEmployeeRepository := infauth.NewRedisEmployeeSessionRepository(rdb)
	sessionManager := authorization.NewSessionManagerImp(sessionEmployeeRepository, time.Duration(cfg.SessionExpiration) * time.Second)

	authenticator := operauth.NewAuthenticator(jwtService, sessionManager)

	return authenticator
}

func NewLogger() *zap.Logger {
	cfg := config.NewConfig()

	return logger.NewLogger(cfg)
}

func NewNewRelicApp() *newrelic.Application {
	cfg := config.NewConfig()

	var app *newrelic.Application
	if cfg.NewRelicLicense != "" {
		var err error
		app, err = newrelic.NewApplication(
			newrelic.ConfigAppName(cfg.LoggerPathToLogFile),
			newrelic.ConfigLicense(cfg.NewRelicLicense),
			newrelic.ConfigDebugLogger(os.Stdout),
		)
		interror.Check(err)
	}

	return app
}