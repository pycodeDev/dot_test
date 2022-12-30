package services

import (
	"context"

	"dot.go/config"
	"dot.go/entities"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/repos"
	"github.com/gofiber/fiber/v2"
)

type ServicePeminjamanImpl struct{}

func NewServicePeminjamanImpl() ServicePeminjaman {
	return &ServicePeminjamanImpl{}
}

func (s ServicePeminjamanImpl) ListPeminjaman(c *fiber.Ctx) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	data, err := repos.NewRepoPeminjamanImpl(db).ListPeminjaman(ctx)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	if len(data) == 0 {
		return model.ReturnError(c, "Data Not Found", rcode.PAGE_NOT_FOUND)
	}

	return model.ReturnSuccess(c, data, "Data Found")
}

func (s ServicePeminjamanImpl) InsertPeminjaman(c *fiber.Ctx, param entities.Peminjaman) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	err := repos.NewRepoPeminjamanImpl(db).InsertPeminjaman(ctx, param)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	return model.ReturnDefault(c, "Success Insert Data", rcode.RESPONSE_SUCCESS)
}
