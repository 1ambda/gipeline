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
	"github.com/a-trium/gipeline/server-gateway/kafka"
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

	kProducer := kafka.NewKafkaProducer(logger, env.Brokers)

	r := mux.NewRouter().StrictSlash(true)
	apiRoute := r.PathPrefix("/api/v1").Subrouter().StrictSlash(true)

	service.RegisterCountryRouter(ctx, kProducer, countryRepo, apiRoute)

	http.Handle("/", r)

	defer func() {
		// Close all resources
		logger.Log("message", "Close all resources")
		kafka.DeleteKafkaProduce(logger, kProducer)
	}()

	logger.Log("error", http.ListenAndServe(env.Port, nil))

	// TODO: graceful shutdown
	// TODO: accessControl
	// TODO: /metrics
	// TODO: globall logging
	// TODO: Number type: int64

}
