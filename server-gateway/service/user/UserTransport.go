package user

import (
	"errors"
	"net/http"

	"github.com/a-trium/gipeline/server-gateway/service/common"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func DecodeUserListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return UserListRequest{}, nil
}

func NewUserListHandler(ctx context.Context, svc UserService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewUserListEndpoint(svc),
		DecodeUserListRequest,
		common.EncodeResponse,
	)
}

func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	u, ok := vars["user"]
	if !ok {
		// TODO error list
		return nil, errors.New("Bad Request for DecodeQueryRequest")
	}

	return UserRequest{User: u}, nil
}

func NewUserHandler(ctx context.Context, svc UserService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewUserEndpoint(svc),
		DecodeUserRequest,
		common.EncodeResponse,
	)
}
