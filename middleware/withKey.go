package middleware

import (
	"dot.go/config"
	"dot.go/model"
	"dot.go/rcode"
	"github.com/gofiber/fiber/v2"
)

func WithKey(c *fiber.Ctx) error {
	k := config.GetMyConfig()
	getApiKey := c.Request().Header.Peek("Api-Key")
	apiKey := string(getApiKey[:])
	if k.APP.API_KEY == apiKey {
		return c.Next()
	} else {
		var bError model.ErrorResponse
		rc := rcode.UNAUTHORIZED
		bError.Status = 0
		bError.ErrorMessage = "Api Key Wrong"
		bError.Rc = rc
		return c.Status(401).JSON(bError)
	}
}
