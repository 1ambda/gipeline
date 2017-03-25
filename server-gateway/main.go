package main

import (
	"net/http"
	"os"

	"github.com/a-trium/gipeline/server-gateway/config"
	"github.com/a-trium/gipeline/server-gateway/service"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
	"github.com/a-trium/gipeline/server-gateway/service/country"
)

func main() {

	flag := config.GetFlag()
	env := config.GetEnvironment()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.NewContext(logger).With(
		"ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller,
	)

	logger.Log(
		"version", flag.Version,
		"gitHash", flag.GitHash,
		"buildTime", flag.BuildTime,
		"started", flag.Started,
		"mode", env.Mode,
		"port", env.Port,
	)

	// Start
	ctx := context.Background()
	countryRepo := country.NewCountryRepository()

	r := mux.NewRouter().StrictSlash(true)
	apiRoute := r.PathPrefix("/api/v1").Subrouter().StrictSlash(true)

	service.RegisterCountryRouter(ctx, countryRepo, apiRoute)

	// TODO: graceful shutdown
	// TODO: accessControl
	// TODO: /metrics
	// TODO: globall logging
	// TODO: Number type: int64

	http.Handle("/", r)
	logger.Log("error", http.ListenAndServe(env.Port, nil))
}
