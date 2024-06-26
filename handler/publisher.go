package handler

import (
	"net/http"
	"strconv"
	"wansanjou/logs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type publisherHandler struct {
	publisherService service.PublisherService
}

func NewPublisherHandler(publisherService service.PublisherService) publisherHandler {
	return publisherHandler{publisherService: publisherService}
}

func (ph publisherHandler) GetPublisherAll(c *fiber.Ctx) error {
	publishers, err := ph.publisherService.GetPublisherAll()
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

	publisher, err := ph.publisherService.GetPublisherByID(id)
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

	response, err := ph.publisherService.CreatePublisher(publisher)
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

	response, err := ph.publisherService.UpdatePublisher(id, publisher)
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

	_, err = ph.publisherService.DeletePublisher(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.SendStatus(http.StatusNoContent)
}
