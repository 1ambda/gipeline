package service

import (
	"github.com/a-trium/gipeline/server-gateway/service/country"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
	"github.com/Shopify/sarama"
)

func RegisterCountryRouter(ctx context.Context, kProducer sarama.SyncProducer, repo country.CountryRepository,
apiRoute *mux.Router) {
	svc := country.NewCountryService(repo)
	route := apiRoute.PathPrefix("/country").Subrouter()

	visitHandler := country.NewCountryVisitHandler(ctx, kProducer, svc)
	route.Handle("/visit", visitHandler).Methods("POST")
}
