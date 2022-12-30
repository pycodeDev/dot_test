package controller

import (
	"dot.go/entities"
	"dot.go/helper"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/services"
	"dot.go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func ListPeminjaman(c *fiber.Ctx) error {
	return services.NewServicePeminjamanImpl().ListPeminjaman(c)
}

func InsertPeminjaman(c *fiber.Ctx) error {
	body := c.Body()
	bodyString := string(body)
	dataJson := gjson.GetMany(bodyString, "id_buku", "created", "expired")
	loop := len(dataJson)
	for i := 0; i < loop; i++ {
		if !dataJson[i].Exists() {
			return model.ReturnError(c, "JSON not valid !!", rcode.PAYLOAD_NOT_VALID)
		}
	}

	id_user, err := utils.GetIDUser(c)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.UNAUTHORIZED)
	}

	param := entities.Peminjaman{
		IdUser:  int32(helper.StringToInt(id_user)),
		IdBuku:  int32(dataJson[0].Int()),
		Status:  1,
		Created: dataJson[1].String(),
		Expired: dataJson[2].String(),
	}

	return services.NewServicePeminjamanImpl().InsertPeminjaman(c, param)
}
