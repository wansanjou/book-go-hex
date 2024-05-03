package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wansanjou/logs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type authorHandler struct {
	authorService service.AuthorService
}

func NewAuthorHandler(authorService service.AuthorService) authorHandler  {
	return authorHandler{authorService: authorService}
}

func (ah authorHandler) GetAuthorAll(c *fiber.Ctx) error  {
	authors , err := ah.authorService.GetAuthorAll()
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Record author not found"})
	}

	return c.JSON(authors)
}

func (ah authorHandler) GetAuthorByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	author, err := ah.authorService.GetAuthorByID(id)
	if err != nil {	
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Author not found"})
	}

	return c.JSON(author)
}

func (ah authorHandler) CreateAuthor(c *fiber.Ctx) error  {
	var author service.AuthorResponse
	if err := c.BodyParser(&author) ; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "Invalid request"})
	}

	response , err := ah.authorService.CreateAuthor(author)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create author"})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (ah authorHandler) UpdateAuthor(c *fiber.Ctx) error  {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var author service.AuthorResponse
	if err := c.BodyParser(&author) ; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "Invalid request"})
	}

	response , err := ah.authorService.UpdateAuthor(id ,author)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update author"})
	}

	return c.JSON(response)
}

func (ah authorHandler) DeleteAuthor(c *fiber.Ctx) error  {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, err = ah.authorService.DeleteAuthor(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.SendStatus(http.StatusNoContent)
}