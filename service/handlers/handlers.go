// Package handlers contains the full set of handler functions and routes
// supported by the http api
package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/tbd54566975/vc-service/did"
	"github.com/tbd54566975/vc-service/framework"
	"github.com/tbd54566975/vc-service/middleware"
	didrouter "github.com/tbd54566975/vc-service/service/routers/did"
)

func API(build string, shutdown chan os.Signal, log *log.Logger, db sql.DB) *framework.Service {
	service := framework.NewService(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))

	readiness := readiness{
		log: log,
	}

	didRouter := didrouter.Handlers{
		DID: did.NewState(log, db),
	}

	service.Handle(http.MethodPost, "/did", didRouter.Create)

	// attach all handlers here
	service.Handle(http.MethodGet, "/health", health)
	service.Handle(http.MethodGet, "/readiness", readiness.handle)

	return service
}
