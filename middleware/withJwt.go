package middleware

import (
	"fmt"
	"time"

	"dot.go/helper"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/services"
	"dot.go/utils"
	"github.com/gofiber/fiber/v2"
)

func WithJwt(c *fiber.Ctx) error {
	funcNow := "WithJwt"
	userAgent := fmt.Sprintf("%s", c.Request().Header.UserAgent())
	dataJwt, err := utils.GetDataFromJwt(c)
	if err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script: failed get data from jwt")
		return c.Status(401).JSON(model.ErrorResponse{
			0,
			rcode.UNAUTHORIZED,
			err.Error(),
		})
	}

	id := fmt.Sprintf("%v", dataJwt["id"])
	token := fmt.Sprintf("%v", dataJwt["token"])
	expired := fmt.Sprintf("%v", dataJwt["expired"])
	time_now := int(time.Now().Unix())
	if time_now > helper.StringToInt(expired) {
		helper.LogError("Token Expired", "func:"+funcNow, "script: Token Expired")
		return c.Status(401).JSON(model.ErrorResponse{
			0,
			rcode.UNAUTHORIZED,
			"Token Expired",
		})
	}

	isValid, err := services.NewServiceUserImpl().UserValidateToken(int32(helper.StringToInt(id)), token, userAgent)
	if err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script: validate token admin")
		return c.Status(401).JSON(model.ErrorResponse{
			0,
			rcode.UNAUTHORIZED,
			err.Error(),
		})
	}

	if !isValid {
		helper.LogError(err.Error(), "func:"+funcNow, "script: token login not valid")
		return c.Status(401).JSON(model.ErrorResponse{
			0,
			rcode.UNAUTHORIZED,
			"Token kamu tidak valid atau expired. Silahkan login ulang kembali !!",
		})
	}

	return c.Next()
}
