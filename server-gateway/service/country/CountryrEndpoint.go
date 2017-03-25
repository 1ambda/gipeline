package country

import (
	"github.com/a-trium/gipeline/server-gateway/service/common"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func NewCountryVisitEndpoint(svc CountryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VisitRequest)

		msg, err := svc.Visit(req.Country)
		res := VisitResponse{
			Message:     msg,
			ErrResponse: *common.NewErrResponse(err),
		}

		return res, nil
	}
}
