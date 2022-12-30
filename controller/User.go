package controller

import (
	"fmt"

	"dot.go/entities"
	"dot.go/helper"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/services"
	"dot.go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func UserLogin(c *fiber.Ctx) error {
	body := c.Body()
	bodyString := string(body)
	dataJson := gjson.GetMany(bodyString, "email", "pass")
	loop := len(dataJson)
	for i := 0; i < loop; i++ {
		if !dataJson[i].Exists() {
			return model.ReturnError(c, "JSON not valid !!", rcode.PAYLOAD_NOT_VALID)
		}
	}

	user := entities.User{
		Email:    dataJson[0].String(),
		Password: dataJson[1].String(),
	}

	ua := fmt.Sprintf("%s", c.Request().Header.UserAgent())

	return services.NewServiceUserImpl().UserLogin(c, user, ua)
}

func UpdateUser(c *fiber.Ctx) error {
	body := c.Body()
	bodyString := string(body)
	dataJson := gjson.GetMany(bodyString, "email", "pass")
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

	user := entities.User{
		Email:    dataJson[0].String(),
		Password: dataJson[1].String(),
		ID:       int32(helper.StringToInt(id_user)),
	}

	return services.NewServiceUserImpl().UserUpdate(c, user)
}
