package api

import (
	"net/http"

	"github.com/Selahattinn/bitaksi/internal/api/response"
	model "github.com/Selahattinn/bitaksi/internal/models"
	"github.com/gofiber/fiber/v2"
)

//parseUserFromRequest parsing body information to user model from request
func parseUserFromRequest(c *fiber.Ctx) (*model.User, error) {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return nil, err

	}
	return &user, nil
}

func (a *API) GetUsers(c *fiber.Ctx) error {
	//Get body information from request
	user, err := parseUserFromRequest(c)
	if err != nil {
		return response.Errorf(c, err, http.StatusInternalServerError, err.Error())
	}

	//Get user from db
	u, err := a.service.GetUserService().GetUser(user)
	if err != nil {
		return response.Errorf(c, err, http.StatusInternalServerError, err.Error())
	}
	return response.Write(c, u)
}

func (a *API) CreateUsers(c *fiber.Ctx) error {
	//Get body information from request
	user, err := parseUserFromRequest(c)
	if err != nil {
		return response.Errorf(c, err, http.StatusInternalServerError, err.Error())
	}

	u, err := a.service.GetUserService().CreateUser(user)
	if err != nil {
		return response.Errorf(c, err, http.StatusInternalServerError, err.Error())
	}

	return response.Write(c, u)
}
