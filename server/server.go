package server

import (
	"context"
	"net/http"

	gokithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/user").Handler(gokithttp.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/user/{id}").Handler(gokithttp.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
