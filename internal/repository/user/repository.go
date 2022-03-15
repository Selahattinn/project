package user

import model "github.com/Selahattinn/bitaksi/internal/models"

type Reader interface {
	GetUser(email string) (*model.User, error)
}

type Writer interface {
	CreateUser(user *model.User) (*model.User, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
