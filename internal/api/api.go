package api

import (
	"github.com/Selahattinn/bitaksi/internal/service"
	"github.com/gofiber/fiber/v2"
)

type API struct {
	JWTSigningKey string
	service       service.Service
}

func NewAPI(a *fiber.App, svc service.Service, SigningKey string) *API {
	api := &API{
		service:       svc,
		JWTSigningKey: SigningKey}
	api.setupRoutes(a)
	return api
}

func (a *API) setupRoutes(app *fiber.App) {
	a.userRoutes(app)
}

// userRoutes defines Users routes
func (a *API) userRoutes(app *fiber.App) {
	route := app.Group("/api/v1/users")
	route.Post("/", a.GetUsers)
	//route.Get("/:id", GetUser)
	route.Put("/", a.CreateUsers)
}
