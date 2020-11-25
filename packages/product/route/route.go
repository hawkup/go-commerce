package route

import "github.com/gofiber/fiber/v2"

func GetProducts(c *fiber.Ctx) error {
	return c.SendString("All Products")
}

func GetProduct(c *fiber.Ctx) error {
	return c.SendString("Single Product")
}

func NewProduct(c *fiber.Ctx) error {
	return c.SendString("New Product")
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("Delete Product")
}
