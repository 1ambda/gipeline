package user

import "github.com/a-trium/gipeline/server-gateway/service/common"

type UserListRequest struct {
}

type UserListResponse struct {
	Users []common.User
	common.ErrResponse
}

type UserRequest struct {
	User string
}

type UserResponse struct {
	User  string `json:"user"`
	Total int    `json:"total"`
	common.ErrResponse
}
