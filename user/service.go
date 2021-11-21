package user

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	logLib "go-service-template/lib/log"
)

const loggerName = "user:service"

type User struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Service interface {
	CreateUser(ctx context.Context, userName string, email string) (User, error)
	GetUser(ctx context.Context, userName string) (User, error)
}

type service struct {
	store  Store
	logger logrus.FieldLogger
}

// NewService initialize a new service
func NewService(logger *logrus.Logger, store Store) Service {
	return &service{
		store:  store,
		logger: logger.WithField(logLib.LoggerKey, loggerName),
	}
}

// CreateUser creates a user profile on storage with given username and password
func (s service) CreateUser(ctx context.Context, username string, password string) (User, error) {
	user, err := s.store.CreateUser(username, password)
	if err != nil {
		// TODO: lhan seek better error package
		return User{}, errors.New("failed to create user")
	}
	return user, nil
}

// GetUser fetch the stored user by given id.
func (s service) GetUser(ctx context.Context, id string) (User, error) {
	user, err := s.store.GetUser(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
