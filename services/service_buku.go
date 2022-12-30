package services

import (
	"dot.go/entities"
	"github.com/gofiber/fiber/v2"
)

type Servicebuku interface {
	ListBuku(c *fiber.Ctx) error
	GetBuku(c *fiber.Ctx, id_buku int) error
	InsertBuku(c *fiber.Ctx, buku entities.Buku) error
	UpdateBuku(c *fiber.Ctx, buku entities.Buku) error
	HapusBuku(c *fiber.Ctx, buku entities.Buku) error
}
