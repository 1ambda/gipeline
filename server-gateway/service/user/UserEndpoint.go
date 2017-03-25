package user

import (
	"github.com/a-trium/gipeline/server-gateway/service/common"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func NewUserListEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(UserListRequest)

		res := UserListResponse{
			Users:       svc.Users(),
			ErrResponse: *common.NewErrResponse(nil),
		}

		return res, nil
	}
}

func NewUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)

		total, err := svc.User(req.User)
		res := UserResponse{
			User:        req.User,
			Total:       total,
			ErrResponse: *common.NewErrResponse(err),
		}

		// TODO: test return error
		return res, nil
	}
}
