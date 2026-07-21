package handler

import (
	"errors"
	"strings"

	"github.com/deblinux123/go-fiber/database"
	"github.com/deblinux123/go-fiber/models"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
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

func GetUserByID(c fiber.Ctx) error {
	var user models.User

	id := c.Params("id")
	result := database.DB.First(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid id.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func UpdateUser(c fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	result := database.DB.First(&user, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	var newUser models.User

	if err := c.Bind().Body(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body request.",
		})
	}

	user.Name = newUser.Name
	user.Email = newUser.Email

	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message": "User updated successfully.",
		"user": fiber.Map{
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func DeleteUser(c fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	result := database.DB.First(&user, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"message": "User deleted successfully.",
	})
}
