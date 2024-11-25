package fiberhandlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/mrqiz/task-9/internal/database"
	"github.com/mrqiz/task-9/internal/models"
)

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	database.DB.Find(&contacts)
	return c.JSON(contacts)
}

func GetContact(c *fiber.Ctx) error {
	var contact models.Contact
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is nan",
		})

	}
	err = database.DB.First(&contact, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}
	return c.JSON(contact)
}

func PostContacts(c *fiber.Ctx) error {
	var contact models.Contact

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "malformed body",
		})
	}

	validate := validator.New()
	if err := validate.Struct(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name and phone are required",
		})
	}

	if err := database.DB.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create contact",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}
