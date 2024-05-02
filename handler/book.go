package handler

import (
	"net/http"
	"strconv"
	"wansanjou/logs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) bookHandler {
	return bookHandler{bookService: bookService}
}

func (bh bookHandler) GetBookAll(c *fiber.Ctx) error {
	books, err := bh.bookService.GetBookAll()
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get books"})
	}

	return c.JSON(books)
}

func (bh bookHandler) GetBookByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	book, err := bh.bookService.GetBookByID(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get book"})
	}

	return c.JSON(book)
}

func (bh bookHandler) CreateBook(c *fiber.Ctx) error {
	var book service.BookResponse
	if err := c.BodyParser(&book); err != nil {
		logs.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := bh.bookService.CreateBook(book)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (bh bookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var book service.BookResponse
	if err := c.BodyParser(&book); err != nil {
		logs.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := bh.bookService.UpdateBook(id, book)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
	}

	return c.JSON(response)
}

func (bh bookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, err = bh.bookService.DeleteBook(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.SendStatus(http.StatusNoContent)
}
