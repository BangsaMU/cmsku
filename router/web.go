package router

import (
	noteRoutes "github.com/BangsaMU/internals/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutesWeb(app *fiber.App) {
	web := app.Group("/", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(web)

	web.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World q!",
		})
	})
	web.Get("/helo", hello)

	// GET http://localhost:8080/hello%20world

	/*
		app.Get("/:value", func(c *fiber.Ctx) error {
			return c.SendString("value: " + c.Params("value"))
			// => Get request with value: hello world
		})


		app.Get("/:param", func(c *fiber.Ctx) error {
			return c.SendString("param: " + c.Params("param"))
		})

		app.Post("/", func(c *fiber.Ctx) error {
			return c.SendString("POST request")
		})
	*/

}

// Handler
func home(c *fiber.Ctx) error {
	return c.SendString("HOME made")
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("I made a ☕ for you!")
}
