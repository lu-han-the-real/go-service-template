package server

import (
	"context"
	"encoding/json"
	"go-service-template/user"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser, GetUser endpoint.Endpoint
}

func MakeEndpoints(s user.Server) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	User user.User `json:"user"`
	Err error `json:"err"`
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeCreateUserEndpoint(svc user.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		userResult, err := svc.CreateUser(ctx, req.Username, req.Password)
		if err != nil {
			return CreateUserResponse{user.User{}, err}, err
		}
		return CreateUserResponse{userResult, nil}, nil
	}
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	UserProfile user.User `json:"userProfile"`
	Err         error     `json:"err"`
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeGetUserEndpoint(svc user.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		localUser, err := svc.GetUser(ctx, req.ID)
		if err != nil {
			return GetUserResponse{UserProfile: user.User{}, Err: err}, err
		}
		return GetUserResponse{localUser, nil}, nil
	}
}

func encodeResponse(
	ctx context.Context,
	w http.ResponseWriter,
	response interface{},
) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetUserRequest
	pathVars := mux.Vars(r)
	request = GetUserRequest{ID: pathVars["id"]}
	return request, nil
}
