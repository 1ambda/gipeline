package country

import (
	"encoding/json"
	"net/http"

	"github.com/a-trium/gipeline/server-gateway/service/common"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/log"
)

func DecodeCountryVisitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req VisitRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func NewCountryVisitHandler(ctx context.Context, logger log.Logger, kProducer sarama.SyncProducer, svc CountryService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewCountryVisitEndpoint(svc, logger, kProducer),
		DecodeCountryVisitRequest,
		common.EncodeResponse,
	)
}
