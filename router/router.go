package router

import (
	"github.com/amusfq/go-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	todo := app.Group("todo")
	todo.Get("/create", handlers.CreateTodoView)
	todo.Get("/:id", handlers.GetTodo)
	todo.Delete("/:id", handlers.DeleteTodo)
	todo.Patch("/:id", handlers.UpdateTodo)
	todo.Get("/", handlers.GetAllTodo)
	todo.Post("/", handlers.CreateTodo)
}
