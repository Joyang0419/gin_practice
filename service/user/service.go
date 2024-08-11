package user

import (
	"context"
	"errors"
	"fmt"

	"gin_practice/service/obj"
)

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrDuplicateUsername = errors.New("username is already exist")
	ErrLoginFailed       = errors.New("login failed")
)

type Service struct {
	Query IQuery
	CMD   ICommand
}

func NewService(query IQuery, CMD ICommand) *Service {
	return &Service{Query: query, CMD: CMD}
}

func (s *Service) Register(ctx context.Context, username, password string) (user *obj.User, err error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("[service][user][Register]%w: empty username or password", ErrInvalidInput)
	}

	isExist, err := s.isUsernameExist(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("[service][user][Register]isUsernameExist err: %w", err)
	}
	if isExist {
		return nil, ErrDuplicateUsername
	}

	user, err = s.CMD.CreateUser(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("[service][user][Register]Repo.CreateUser err: %w", err)

	}

	return user, nil
}

func (s *Service) isUsernameExist(ctx context.Context, username string) (bool, error) {
	filter := FilterOfUser{
		Username: &username,
	}
	user, err := s.Query.User(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("[service][user][isUsernameExist]Repo.User err: %w", err)
	}
	if user == nil {
		return false, nil
	}

	return true, nil
}

func (s *Service) Login(ctx context.Context, username, password string) (user *obj.User, err error) {
	filter := FilterOfUser{
		Username: &username,
		Password: &password,
	}
	user, err = s.Query.User(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("[service][user][Login]Repo.User err: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("[service][user][Login]%w: %s", ErrLoginFailed, "username or password is incorrect")
	}

	return user, nil
}
