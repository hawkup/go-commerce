package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spiritsquare/go-commerce/packages/product/route"
	"github.com/spiritsquare/go-commerce/packages/shared/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/v1/product", route.GetProducts)
	app.Get("/api/v1/product/:id", route.GetProduct)
	app.Post("/api/v1/product", route.NewProduct)
	app.Delete("/api/v1/product/:id", route.DeleteProduct)
}

func main() {
	app := fiber.New()
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
