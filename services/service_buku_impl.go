package services

import (
	"context"
	"encoding/json"
	"time"

	"dot.go/config"
	"dot.go/entities"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/repos"
	"github.com/gofiber/fiber/v2"
)

type ServiceBukuImpl struct{}

func NewServiceBukuImpl() Servicebuku {
	return &ServiceBukuImpl{}
}

func (s ServiceBukuImpl) ListBuku(c *fiber.Ctx) error {
	ctx := context.Background()
	db := config.GormConnectWrite()
	var out []entities.Buku

	rd := config.Redislocal()
	defer rd.Close()
	key := rcode.LIST_BUKU

	res, _ := rd.Get(ctx, key).Result()
	if res != "" {
		json.Unmarshal([]byte(res), &out)
		return c.JSON(model.SuccessResponse{
			Status:    1,
			Rc:        rcode.RESPONSE_SUCCESS,
			Message:   "Data Found",
			Data:      out,
			TimeStamp: time.Now().Unix(),
		})
	}

	data, err := repos.NewRepoBukuImpl(db).ListBuku(ctx)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	if len(data) == 0 {
		return model.ReturnError(c, "Data Not Found", rcode.PAGE_NOT_FOUND)
	}

	save, _ := json.Marshal(data)

	rd.SetEX(ctx, key, save, time.Minute*10)

	return model.ReturnSuccess(c, data, "Data Found")
}

func (s ServiceBukuImpl) GetBuku(c *fiber.Ctx, id_buku int) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	data, err := repos.NewRepoBukuImpl(db).GetBuku(ctx, id_buku)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	if data.ID == 0 {
		return model.ReturnError(c, "Data Not Found", rcode.PAGE_NOT_FOUND)
	}

	return model.ReturnSuccess(c, data, "Data Found")
}

func (s ServiceBukuImpl) InsertBuku(c *fiber.Ctx, buku entities.Buku) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	rd := config.Redislocal()
	defer rd.Close()
	key := rcode.LIST_BUKU

	rd.Del(ctx, key)

	err := repos.NewRepoBukuImpl(db).InsertBuku(ctx, buku)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	return model.ReturnDefault(c, "Data Success Insert", rcode.RESPONSE_SUCCESS)
}

func (s ServiceBukuImpl) UpdateBuku(c *fiber.Ctx, buku entities.Buku) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	rd := config.Redislocal()
	defer rd.Close()
	key := rcode.LIST_BUKU

	rd.Del(ctx, key)

	err := repos.NewRepoBukuImpl(db).UpdateBuku(ctx, buku)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	return model.ReturnDefault(c, "Data Success Update", rcode.RESPONSE_SUCCESS)
}

func (s ServiceBukuImpl) HapusBuku(c *fiber.Ctx, buku entities.Buku) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	rd := config.Redislocal()
	defer rd.Close()
	key := rcode.LIST_BUKU

	rd.Del(ctx, key)

	err := repos.NewRepoBukuImpl(db).HapusBuku(ctx, buku)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	return model.ReturnDefault(c, "Data Success Delete", rcode.RESPONSE_SUCCESS)
}
