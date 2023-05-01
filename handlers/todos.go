package handlers

import (
	"math"
	"strconv"
	"github.com/amusfq/go-fiber/database"
	"github.com/amusfq/go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodo(c *fiber.Ctx) error {
	data := []models.Todo{}
	search := c.Query("search")
    perPage, err := strconv.Atoi(c.Query("perPage", "10"))
	if err != nil {
		return c.SendString("perPage harus angka")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("page harus angka")
	}

    offset := (page - 1) * perPage
	result := database.DB.Db

	if search != "" {
		result = result.Where("title like ?", "%"+search+"%")
	}
	result.Offset(offset).Limit(perPage).Find(&data)

	var total int64

	if search != "" {
		database.DB.Db.Where("title like ?", "%"+search+"%").Count(&total)
	} else {
		database.DB.Db.Model(&models.Todo{}).Count(&total)
	}
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return c.Render("todo/index", fiber.Map{
		"PageTitle":       "List Todo",
		"PageDescription": "List of Todo",
		"Data":            data,
		"TotalData": total,
		"Page": int(page),
		"TotalPages": totalPages,
		"PerPage": perPage,
	})
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	data := models.Todo{}
	result := database.DB.Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return c.Status(404).SendString("Not Found")
	}
	return c.Render("todo/view", fiber.Map{
		"Data": data,
	})
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

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(models.Todo)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	if data.Title == "" {
		return c.Status(500).JSON(fiber.Map{"Message": "Title is required"})
	}
	result := database.DB.Db.Model(&data).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": result.Error,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"Message": "Success",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	data := models.Todo{}
	database.DB.Db.Delete(&data, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
