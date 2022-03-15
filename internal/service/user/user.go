package user

import (
	"errors"

	model "github.com/Selahattinn/bitaksi/internal/models"
	"github.com/Selahattinn/bitaksi/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	repository repository.Repository
}

var (
	ErrorUserNotFound = errors.New("User not found")
)

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) GetUser(user *model.User) (*model.User, error) {
	//Get user from db
	u, err := s.repository.GetUserRepository().GetUser(user.Email)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func (s *Service) CreateUser(user *model.User) (*model.User, error) {
	//Get user from db
	u, err := s.repository.GetUserRepository().CreateUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}
