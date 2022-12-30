package routes

import (
	"dot.go/controller"
	"dot.go/middleware"
	"dot.go/rcode"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", controller.Index)

	user := api.Group("/user")
	user.Post("/login", controller.UserLogin)
	user.Patch("/update", middleware.WithKey, middleware.WithJwt, controller.UpdateUser)

	buku := api.Group("/buku")
	buku.Get("/list", middleware.WithKey, middleware.WithJwt, controller.ListBuku)
	buku.Post("/insert", middleware.WithKey, middleware.WithJwt, controller.InsertBuku)
	buku.Put("/update", middleware.WithKey, middleware.WithJwt, controller.UpdateBuku)
	buku.Delete("/delete", middleware.WithKey, middleware.WithJwt, controller.HapusBuku)
	buku.Get("/:id", middleware.WithKey, middleware.WithJwt, controller.GetBuku)

	peminjaman := api.Group("/peminjaman")
	peminjaman.Get("/list", middleware.WithKey, middleware.WithJwt, controller.ListPeminjaman)
	peminjaman.Post("/create", middleware.WithKey, middleware.WithJwt, controller.InsertPeminjaman)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(rcode.PAGE_NOT_FOUND).JSON(fiber.Map{"status": 0, "rc": rcode.PAGE_NOT_FOUND, "error_msg": "Page Not Found!"}) // => 404 "Not Found"
	})
}
