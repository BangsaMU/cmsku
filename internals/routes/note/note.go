package noteRoutes

import (
	noteHandler "github.com/BangsaMU/internals/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	// note.Post("/", noteHandler.CreateNotes)

	// Read all Notes
	note.Get("/", noteHandler.GetNotes)
	// // Read one Note
	note.Get("/:noteId", noteHandler.GetNote)

	// Update one Note
	note.Put("/:noteId", noteHandler.UpdateNote)
	// Delete one Note
	note.Delete("/:noteId", noteHandler.DeleteNote)
	note.Get("/delete/:noteId", noteHandler.DeleteNote)
}
