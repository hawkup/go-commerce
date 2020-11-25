package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/spiritsquare/go-commerce/packages/order/route"
	"github.com/spiritsquare/go-commerce/packages/shared/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/v1/order", route.GetOrders)
	app.Get("/api/v1/order/:id", route.GetOrder)
	app.Post("/api/v1/order", route.NewOrder)
	app.Delete("/api/v1/order/:id", route.DeleteOrder)
}

func main() {
	app := fiber.New()
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
