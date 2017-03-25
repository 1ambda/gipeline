package number

import (
	"encoding/json"
	"net/http"

	"github.com/a-trium/gipeline/server-gateway/service/common"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func DecodeInsertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func NewInsertHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		common.EncodeResponse,
	)
}

func NewTotalHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewTotalEndpoint(svc),
		common.DecodeEmptyRequest,
		common.EncodeResponse,
	)
}
