package controller

import (
	"dot.go/model"
	"dot.go/rcode"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return model.ReturnDefault(c, "Welcome To Api v1 DOT", rcode.RESPONSE_SUCCESS)
}
