package routes

import (
	controllers "gomvc02/Controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	// Route setup code would go here
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:56874", // หรือ "*" ถ้าจะให้ทุก origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running!")
	})
	app.Post("/Api/register", controllers.RegisterUserLogin)
	app.Get("/Api/testconnect", controllers.StatusconnectDB)
	app.Post("/Api/Loginuser", controllers.Loginuser)
}
