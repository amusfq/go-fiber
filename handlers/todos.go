package handlers

import (
	"github.com/amusfq/go-fiber/database"
	"github.com/amusfq/go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodo(c *fiber.Ctx) error {
	data := []models.Todo{}
	database.DB.Db.Find(&data)
	return c.Render("todo/index", fiber.Map{
		"PageTitle":       "List Todo",
		"PageDescription": "List of Todo",
		"Data":            data,
	})
}

func GetTodo(c *fiber.Ctx) error {
	return c.SendString("Div Rhino Trivia App!")
}

func CreateTodoView(c *fiber.Ctx) error {
	return c.Render("todo/create", fiber.Map{
		"PageTitle":       "Create Todo",
		"PageDescription": "Create new Todo",
	})
}

func CreateTodo(c *fiber.Ctx) error {
	data := new(models.Todo)
	if err := c.BodyParser(data); err != nil {
		return c.Render("todo/error", fiber.Map{"Error": err.Error()})
	}
	if data.Title == "" {
		return c.Render("todo/error", fiber.Map{"Error": "Title is required"})
	}

	database.DB.Db.Create(&data)
	return c.Render("todo/success", fiber.Map{})
}
