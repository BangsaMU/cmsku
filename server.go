package main

import (
	"github.com/BangsaMU/database"
	apiRouter "github.com/BangsaMU/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	// "gorm.io/gorm/logger"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./internals/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Connect to the Database
	database.ConnectDB()

	// Setup the router api note v1
	apiRouter.SetupRoutesNote(app)
	apiRouter.SetupRoutesNoteV2(app)
	apiRouter.SetupRoutesWeb(app)

	//Assets
	app.Static("/", "./public")

	//API
	/*
		app.Get("/api/posts", func(c *fiber.Ctx) error {
			posts := getPosts() // your logic
			if len(posts) == 0 {
				return c.Status(404).JSON(&fiber.Map{
					"success": false,
					"error":   "There are no posts!",
				})
			}
			return c.JSON(&fiber.Map{
				"success": true,
				"posts":   posts,
			})
		})
	*/

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}
