package main

import (
	"log"

	"github.com/deblinux123/go-fiber/database"
	"github.com/deblinux123/go-fiber/handler"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// connecting to database
	database.Connect()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hello Fiber!")
	})

	// add new router
	app.Get("/about", func(c fiber.Ctx) error {
		return c.SendString("This is about page and nothing else.")
	})

	app.Get("/wellcome", func(c fiber.Ctx) error {
		return c.SendString("Welcome to my page")
	})

	app.Get("/contact", func(c fiber.Ctx) error {
		return c.SendString("Contact us")
	})

	// return json file
	app.Get("/user", handler.GetUsers)

	// post request
	app.Post("/signup", handler.SignUp)

	// get user by id
	app.Get("/user/:id", handler.GetUserByID)

	// update the user information
	app.Put("/user/:id", handler.UpdateUser)

	// delete the user
	app.Delete("/user/:id", handler.DeleteUser)

	log.Fatal(app.Listen(":3000"))
}
