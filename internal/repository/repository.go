package repository

import "github.com/Selahattinn/bitaksi/internal/repository/user"

type Repository interface {
	Shutdown()
	GetUserRepository() user.Repository
}
