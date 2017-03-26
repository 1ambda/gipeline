package country

import (
	"github.com/a-trium/gipeline/server-gateway/service/common"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/log"
	"fmt"
)

func NewCountryVisitEndpoint(svc CountryService, logger log.Logger, kProducer sarama.SyncProducer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VisitRequest)

		msg, err := svc.Visit(req.Country)

		// 1. send message to kafka
		_, _, kErr := kProducer.SendMessage(&sarama.ProducerMessage{
			Topic: "country",
			Value: sarama.StringEncoder(fmt.Sprintf("{\"country\": %s}", req.Country)),
		})

		if kErr != nil {
			logger.Log("error", kErr)
			return nil, kErr
		}

		// 2. persist
		res := VisitResponse{
			Message:     msg,
			ErrResponse: *common.NewErrResponse(err),
		}

		return res, nil
	}
}
