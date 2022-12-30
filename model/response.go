package model

import (
	"time"

	"dot.go/rcode"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status       int    `json:"status"`
	Rc           int    `json:"rc"`
	ErrorMessage string `json:"error_msg"`
}

type ErrorResponseWithData struct {
	Status int      `json:"status"`
	Rc     int      `json:"rc"`
	Error  []string `json:"error"`
}

type ResponseDefault struct {
	Status  int    `json:"status"`
	Rc      int    `json:"rc"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status    int         `json:"status"`
	Rc        int         `json:"rc"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	TimeStamp int64       `json:"ts"`
}

func ReturnError(c *fiber.Ctx, msg string, status int) error {
	return c.Status(status).JSON(ErrorResponse{
		0,
		status,
		msg,
	})
}

func ReturnErrorWithData(c *fiber.Ctx, msg []string, status int) error {
	return c.Status(status).JSON(ErrorResponseWithData{
		0,
		status,
		msg,
	})
}

func ReturnDefault(c *fiber.Ctx, msg string, status int) error {
	return c.Status(status).JSON(ResponseDefault{
		1,
		status,
		msg,
	})
}

func ReturnSuccess(c *fiber.Ctx, data interface{}, msg string) error {
	return c.Status(200).JSON(SuccessResponse{
		1,
		rcode.RESPONSE_SUCCESS,
		msg,
		data,
		time.Now().Unix(),
	})
}
