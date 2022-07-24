package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
  Id string
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "camilo",
		Lastname:  "Davila",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	
	if err := c.BodyParser(&user); err != nil {
		return err
	}

  user.Id = uuid.NewString()
  
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()
	
	// Middlewares
	app.Use(logger.New())
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world! in go")
	})

  userGroup := app.Group("/users")
	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)
	
	app.Listen(":3000")
}
