package services

import (
	"dot.go/entities"
	"github.com/gofiber/fiber/v2"
)

type ServiceUser interface {
	UserLogin(c *fiber.Ctx, user entities.User, ua string) error
	UserValidateToken(id_user int32, token string, ua string) (bool, error)
	UserUpdate(c *fiber.Ctx, user entities.User) error
}
