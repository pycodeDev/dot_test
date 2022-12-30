package services

import (
	"dot.go/entities"
	"github.com/gofiber/fiber/v2"
)

type ServicePeminjaman interface {
	ListPeminjaman(c *fiber.Ctx) error
	InsertPeminjaman(c *fiber.Ctx, param entities.Peminjaman) error
}
