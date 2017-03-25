package number

import (
	"github.com/a-trium/gipeline/server-gateway/service/common"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func NewInsertEndpoint(svc NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertRequest)

		msg, err := svc.Insert(req.User, req.Number)
		res := InsertResponse{
			Message:     msg,
			ErrResponse: *common.NewErrResponse(err),
		}

		return res, nil
	}
}

func NewTotalEndpoint(svc NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res := TotalResponse{
			Total:       svc.Total(),
			ErrResponse: *common.NewErrResponse(nil),
		}

		return res, nil
	}
}
