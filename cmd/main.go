package main

import (
	"github.com/amusfq/go-fiber/database"
	"github.com/amusfq/go-fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layouts/main",
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
