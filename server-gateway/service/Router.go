package service

import (
	"github.com/a-trium/gipeline/server-gateway/service/number"
	"github.com/a-trium/gipeline/server-gateway/service/user"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

func RegisterNumberRouter(ctx context.Context, repo number.NumberRepository,
	apiRoute *mux.Router) {
	svc := number.NewNumberService(repo)
	route := apiRoute.PathPrefix("/number").Subrouter()

	insertHandler := number.NewInsertHandler(ctx, svc)
	route.Handle("/insert", insertHandler).Methods("POST")

	totalHandler := number.NewTotalHandler(ctx, svc)
	route.Handle("/total", totalHandler).Methods("GET")
}

func RegisterUserRouter(ctx context.Context, repo number.NumberRepository,
	apiRoute *mux.Router) {
	svc := user.NewUserService(repo)
	route := apiRoute.PathPrefix("/user").Subrouter()

	usersHandler := user.NewUserListHandler(ctx, svc)
	// workaround: https://github.com/gorilla/mux/issues/31
	apiRoute.Handle("/user", usersHandler).Methods("GET")

	userHandler := user.NewUserHandler(ctx, svc)
	route.Handle("/{user}", userHandler).Methods("GET")
}
