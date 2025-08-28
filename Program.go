package main

import (
	configs "gomvc02/Configs"
	routes "gomvc02/Routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectDatabase()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3001"))
}
