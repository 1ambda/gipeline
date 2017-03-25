package country

import "github.com/a-trium/gipeline/server-gateway/service/common"

type VisitRequest struct {
	Country string `json:"country"`
}

type VisitResponse struct {
	Message string `json:"message"`
	common.ErrResponse
}
