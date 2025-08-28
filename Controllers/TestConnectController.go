package controllers

import (
	configs "gomvc02/Configs"

	"github.com/gofiber/fiber/v2"
)

func StatusconnectDB(c *fiber.Ctx) error {

	conb := configs.Testconnect()
	if conb {
		return c.Status(200).JSON(fiber.Map{
			"status": "Database connected"})
	} else {
		return c.Status(500).JSON(fiber.Map{
			"status": "Database not connected"})
	}
}
