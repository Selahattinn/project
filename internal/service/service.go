package service

import "github.com/Selahattinn/bitaksi/internal/service/user"

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	Shutdown()
}
