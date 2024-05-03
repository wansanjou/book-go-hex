package handler

import (
	"net/http"
	"strconv"
	"wansanjou/logs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type publisherHandler struct {
	publisher_service service.PublisherService
}

func NewPublisherHandler(publisher_service service.PublisherService) publisherHandler {
	return publisherHandler{publisher_service: publisher_service}
}

func (ph publisherHandler) GetPublisherAll(c *fiber.Ctx) error {
	publishers, err := ph.publisher_service.GetPublisherAll()
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get publishers"})
	}

	return c.JSON(publishers)
}

func (ph publisherHandler) GetPublisherByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	publisher, err := ph.publisher_service.GetPublisherByID(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get publisher"})
	}

	return c.JSON(publisher)
}

func (ph publisherHandler) CreatePublisher(c *fiber.Ctx) error {
	var publisher service.PublisherResponse
	if err := c.BodyParser(&publisher); err != nil {
		logs.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := ph.publisher_service.CreatePublisher(publisher)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create publisher"})
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (ph publisherHandler) UpdatePublisher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var publisher service.PublisherResponse
	if err := c.BodyParser(&publisher); err != nil {
		logs.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := ph.publisher_service.UpdatePublisher(id, publisher)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update publisher"})
	}

	return c.JSON(response)
}

func (ph publisherHandler) DeletePublisher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, err = ph.publisher_service.DeletePublisher(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.SendStatus(http.StatusNoContent)
}
