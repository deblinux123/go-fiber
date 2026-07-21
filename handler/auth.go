package handler

import (
	"strings"

	"github.com/deblinux123/go-fiber/database"
	"github.com/deblinux123/go-fiber/models"
	"github.com/gofiber/fiber/v3"
)

func SignUp(c fiber.Ctx) error {
	var user models.User

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body.",
		})
	}

	if strings.TrimSpace(user.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	if strings.TrimSpace(user.Email) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	if len(user.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password must be at leats 6 charachters or more",
		})
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully.",
		"user": fiber.Map{
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func GetUsers(c fiber.Ctx) error {
	var users []models.User

	result := database.DB.Find(&users)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": users,
	})
}
