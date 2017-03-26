package service

import (
	"github.com/a-trium/gipeline/server-gateway/service/country"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/log"
)

func RegisterCountryRouter(ctx context.Context, logger log.Logger, kProducer sarama.SyncProducer, repo country.CountryRepository,
apiRoute *mux.Router) {
	svc := country.NewCountryService(repo)
	route := apiRoute.PathPrefix("/country").Subrouter()

	visitHandler := country.NewCountryVisitHandler(ctx, logger, kProducer, svc)
	route.Handle("/visit", visitHandler).Methods("POST")
}
