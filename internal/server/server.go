package server

import (
	"net/http"
	"os"

	"github.com/Selahattinn/bitaksi/internal/api"
	"github.com/Selahattinn/bitaksi/internal/repository"
	"github.com/Selahattinn/bitaksi/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string `yaml:"host"`
	// For HTTPS
	//CertFile      string `yaml:"cert_file"`
	//KeyFile       string `yaml:"key_file"`
	SigningKey string                  `yaml:"signing_key"`
	DB         *repository.MongoConfig `yaml:"database"`
	Service    *service.Config         `yaml:"service"`
}

// Instance represents an instance of the server
type Instance struct {
	Config *Config
	Fiber  *fiber.App
	API    *api.API
}

// NewInstance returns an new instance of our server
func NewInstance(cfg *Config) *Instance {
	return &Instance{
		Config: cfg,
		Fiber:  fiber.New(),
	}
}

// Start starts the server
func (i *Instance) Start() {

	repository, err := repository.NewMongoRepository(i.Config.DB)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start MongoDB")
	}
	service, err := service.NewProvider(i.Config.Service, repository)

	if err != nil {
		logrus.WithError(err).Fatal("Failed to start service")
	}
	i.API = api.NewAPI(i.Fiber, service, i.Config.SigningKey)

	err = i.Fiber.Listen(i.Config.ListenAddress)
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("Fiber Server stopped unexpected")
		i.Shutdown()
	} else {
		logrus.WithError(err).Info("Fiber Server stopped")
	}
}

// Shutdown stops the server gracefully
func (i *Instance) Shutdown() {
	err := i.Fiber.Shutdown()
	if err != nil {
		logrus.WithError(err).Error("Failed to shutdown Fiber server gracefully")
		os.Exit(1)
	}

	logrus.Info("Shutdown Fiber server...")
	os.Exit(0)
}
