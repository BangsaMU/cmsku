package router

import (
	noteRoutes "github.com/BangsaMU/internals/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutesNote(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(api)

	/*
		handler := func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		}
			api := app.Group("/api", logger.New())

			// Setup the Node Routes
			noteRoutes.SetupNoteRoutes(api)

			v1 := api.Group("/v1")   // /api/v1
			v1.Get("/list", handler) // /api/v1/list
			v1.Get("/user", handler) // /api/v1/user

			v2 := api.Group("/v2")   // /api/v2
			v2.Get("/list", handler) // /api/v2/list
			v2.Get("/user", handler) // /api/v2/user
	*/
}

func SetupRoutesNoteV2(app *fiber.App) {
	apiv2 := app.Group("/api/v2", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(apiv2)
}
