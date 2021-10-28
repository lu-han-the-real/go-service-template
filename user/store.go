package user

import (
	"errors"
	"github.com/lithammer/shortuuid/v3"
)

// Store interface to provide User's Data's CRUD interaction via database storage
type Store interface {
	CreateUser(username,password string) (User, error)
	GetUser(id string) (User, error)
}

type localStore struct {
	UserProfiles map[string]User
}

func NewStore() Store{
	userProfiles := make(map[string]User, 0)
	return &localStore{UserProfiles: userProfiles}
}

func (l *localStore) CreateUser(username, password string) (User, error){
	if username == "" {
		return User{}, errors.New("user name should not be empty")
	}

	if password == "" {
		return User{}, errors.New("password should not be empty")
	}

	userID := shortuuid.New()
	user := User{
		ID:       userID,
		UserName: username,
		Password: password,
	}
	l.UserProfiles[userID] = user

	return user, nil
}

func (l *localStore) GetUser (id string) (User, error){
	user, ok := l.UserProfiles[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}