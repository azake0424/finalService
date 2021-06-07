package handlers

import (
	"finalService/foundation/web"
	"log"
	"net/http"
	"os"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown)

	check := check{
		logger: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
