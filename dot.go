package main

import (
	"log"

	"dot.go/config"
	"dot.go/helper"
	"dot.go/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func testLog() {
}

func main() {
	testLog()
	runServer()
}

func runServer() {
	k := config.GetMyConfig()
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			if code == fiber.StatusInternalServerError {
				helper.LogError(err.Error(), "funcNow:Main")
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": 0, "error_msg": err.Error()})
			}
			return nil
		},
	})

	app.Use(cors.New())
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	))
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		c.Set("X-Powered-By", k.APP.APP_NAME)
		c.Set("X-Version", k.APP.VERSION)
		c.Set("X-Software", "Golang")
		return c.Next()
	})

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":" + k.APP.MY_PORT))
}
