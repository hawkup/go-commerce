package route

import "github.com/gofiber/fiber/v2"

func GetOrders(c *fiber.Ctx) error {
	return c.SendString("All Orders")
}

func GetOrder(c *fiber.Ctx) error {
	return c.SendString("Single Order")
}

func NewOrder(c *fiber.Ctx) error {
	return c.SendString("New Order")
}

func DeleteOrder(c *fiber.Ctx) error {
	return c.SendString("Delete Order")
}
