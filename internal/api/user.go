package api

import (
	"fmt"

	"github.com/Selahattinn/bitaksi/internal/api/response"
	"github.com/gofiber/fiber/v2"
)

func (a *API) GetUsers(c *fiber.Ctx) error {
	fmt.Println("asdasd")
	return response.Write(c, "asdasdas")
}
