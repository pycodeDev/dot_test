package controller

import (
	"dot.go/entities"
	"dot.go/helper"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/services"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func ListBuku(c *fiber.Ctx) error {
	return services.NewServiceBukuImpl().ListBuku(c)
}

func GetBuku(c *fiber.Ctx) error {
	id := helper.StringToInt(c.Params("id"))
	return services.NewServiceBukuImpl().GetBuku(c, id)
}

func InsertBuku(c *fiber.Ctx) error {
	body := c.Body()
	bodyString := string(body)
	dataJson := gjson.GetMany(bodyString, "nama", "status")
	loop := len(dataJson)
	for i := 0; i < loop; i++ {
		if !dataJson[i].Exists() {
			return model.ReturnError(c, "JSON not valid !!", rcode.PAYLOAD_NOT_VALID)
		}
	}

	param := entities.Buku{
		Nama:      dataJson[0].String(),
		Status:    int32(dataJson[1].Int()),
		CreatedAt: helper.GetCurrentDateTime(),
	}

	return services.NewServiceBukuImpl().InsertBuku(c, param)
}

func UpdateBuku(c *fiber.Ctx) error {
	body := c.Body()
	bodyString := string(body)
	dataJson := gjson.GetMany(bodyString, "nama", "status", "id")
	loop := len(dataJson)
	for i := 0; i < loop; i++ {
		if !dataJson[i].Exists() {
			return model.ReturnError(c, "JSON not valid !!", rcode.PAYLOAD_NOT_VALID)
		}
	}

	param := entities.Buku{
		Nama:   dataJson[0].String(),
		Status: int32(dataJson[1].Int()),
		ID:     int32(dataJson[2].Int()),
	}

	return services.NewServiceBukuImpl().UpdateBuku(c, param)
}

func HapusBuku(c *fiber.Ctx) error {
	body := c.Body()
	bodyString := string(body)
	dataJson := gjson.GetMany(bodyString, "id")
	loop := len(dataJson)
	for i := 0; i < loop; i++ {
		if !dataJson[i].Exists() {
			return model.ReturnError(c, "JSON not valid !!", rcode.PAYLOAD_NOT_VALID)
		}
	}

	param := entities.Buku{
		ID: int32(dataJson[0].Int()),
	}

	return services.NewServiceBukuImpl().HapusBuku(c, param)
}
