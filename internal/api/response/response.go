package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Error   bool   `json:"error"`
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}

// Errorf return an new error response
func Errorf(a *fiber.Ctx, err error, code int, message string) error {
	logrus.WithFields(logrus.Fields{
		"host":       string(a.Request().Host()),
		"address":    a.Context().RemoteAddr().String(),
		"method":     a.Method(),
		"requestURI": a.Request().URI().String(),
		"proto":      a.Protocol(),
		"useragent":  string(a.Context().UserAgent()),
	}).WithError(err).Debug(message)

	errorMessage := Error{
		Error:   true,
		Code:    code,
		Message: message,
	}
	return a.Status(code).JSON(errorMessage)
}

// Write return a new json response
func Write(a *fiber.Ctx, data interface{}) error {
	logrus.WithFields(logrus.Fields{
		"host":       string(a.Request().Host()),
		"address":    a.Context().RemoteAddr().String(),
		"method":     a.Method(),
		"requestURI": a.Request().URI().String(),
		"proto":      a.Protocol(),
		"useragent":  string(a.Context().UserAgent()),
	}).Debug(data)

	return a.Status(200).JSON(data)
}
