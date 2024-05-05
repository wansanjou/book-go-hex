package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wansanjou/logs"
	"wansanjou/service"

	"github.com/gofiber/fiber/v2"
)

type userHandle struct {
	userService service.UserService
}

func NewUserService(userService service.UserService) userHandle  {
	return userHandle{userService: userService}
}

func (uh userHandle) GetUserAll(c *fiber.Ctx) error  {
	users , err := uh.userService.GetUserAll()
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Record user not found"})
	}

	return c.JSON(users)
}

func (uh userHandle) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	user, err := uh.userService.GetUserByID(id)
	if err != nil {	
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (uh userHandle) CreateUser(c *fiber.Ctx) error  {
	var user service.UserResponse
	if err := c.BodyParser(&user) ; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "Invalid request"})
	}

	response , err := uh.userService.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}


func (uh userHandle) UpdateUser(c *fiber.Ctx) error  {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var user service.UserResponse
	if err := c.BodyParser(&user) ; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "Invalid request"})
	}

	response , err := uh.userService.UpdateUser(id ,user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(response)
}


func (uh userHandle) DeleteUser(c *fiber.Ctx) error  {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, err = uh.userService.DeleteUser(id)
	if err != nil {
		logs.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(http.StatusNoContent)
}

func (uh userHandle) LoginUser(c *fiber.Ctx) error {

	var user service.UserResponse
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email and password are required",
		})
	}

	response, err := uh.userService.LoginUser(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
