package controllers

import (
	configs "gomvc02/Configs"
	models "gomvc02/Models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("mysecretkeytestgomvc123456")

func RegisterUserLogin(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON"})

	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	if result := configs.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not create user"})
	}
	return c.Status(201).JSON(user)

}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

func Loginuser(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON"})
	}

	var user models.User
	if result := configs.DB.Where("username = ?", input.Username).First(&user); result.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid username or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid username or password"})
	}

	// Token generation logic would go here

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"token":   "testtoken123456",
	})
}
