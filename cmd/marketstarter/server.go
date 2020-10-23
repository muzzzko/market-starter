package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"market-starter/internal/dependencymgr"
	interror "market-starter/internal/error"
	"market-starter/internal/error/grapherr"
	"market-starter/internal/graphql/graph/generated"
	"market-starter/internal/logger"
	"market-starter/internal/operation/authorization"
	intgin "market-starter/internal/operation/gin"
	intnewrelic "market-starter/internal/operation/newrelic"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	cfg := generated.Config{Resolvers: dependencymgr.NewResolver()}

	authenticator := dependencymgr.NewAuthenticator()

	cfg.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		employeeSession := authenticator.AuthenticateAndReturnEmployeeSession(ctx)
		if employeeSession != nil {
			ctx := authorization.AddEmpoyeeSessionToContext(ctx, employeeSession)
			return next(ctx)
		}

		return nil, grapherr.Serialize(interror.NewInvalidCredential())
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	dependencymgr.NewLogger()
	app := dependencymgr.NewNewRelicApp()

	r := gin.Default()

	r.Use(intgin.GinContextToContextMiddleware())
	r.Use(logger.AddLoggerToContextMiddleware())
	r.Use(nrgin.Middleware(app))
	r.Use(intnewrelic.Middleware())
	r.Use(authorization.AddTokenToContextMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
