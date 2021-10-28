package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log/logrus"
)

type User struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Server interface {
	CreateUser(ctx context.Context, userName string, email string) (User, error)
	GetUser(ctx context.Context, userName string) (User, error)
}

type server struct {
	store  Store
	logger logrus.Logger
}

// NewServer initialize a new server
func NewServer(logger logrus.Logger, store Store) Server {
	return &server{
		store:  store,
		logger: logger,
	}
}

// CreateUser creates a user profile on storage with given username and password
func (s server) CreateUser(ctx context.Context, username string, password string) (User, error) {
	user, err := s.store.CreateUser(username, password)
	if err != nil {
		// TODO: lhan seek better error package
		return User{}, errors.New("failed to create user")
	}
	return user, nil
}

// GetUser fetch the stored user by given id.
func (s server) GetUser(ctx context.Context, id string) (User, error) {
	user, err := s.store.GetUser(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
