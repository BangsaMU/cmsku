package noteHandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/BangsaMU/database"
	"github.com/BangsaMU/internals/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	// find all notes in the database
	// db.Find(&notes)
	db.Raw("SELECT * FROM notes").Scan(&notes)

	result := db.Find(&notes)
	err := result.Error
	ttl := result.RowsAffected
	msg1 := "total data "
	msg := fmt.Sprintf("%s%d", msg1, ttl)
	errors.Is(err, gorm.ErrRecordNotFound)

	log.Println("error", err)
	log.Println(notes)

	// If no note is present return an error
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": msg, "data": notes})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// result := db.First(&note)
	// result.RowsAffected // returns count of records found
	// result.Error        // returns error or nil

	// Find the note with the given Id
	// db.Find(&note, "id = ?", id)
	db.Where("id = ?", id).Find(&note)
	// db.Raw("SELECT * FROM notes WHERE id = ?", id).Scan(&note)

	result := db.Find(&note, id)
	err := result.Error
	ttl := result.RowsAffected
	msg1 := "total data "
	msg := fmt.Sprintf("%s%d", msg1, ttl)
	errors.Is(err, gorm.ErrRecordNotFound)

	log.Println("RowsAffected", ttl)
	log.Println("error", err)
	log.Println(note)

	// If no such note present return an error
	if ttl == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Return the note with the Id
	return c.JSON(fiber.Map{"status": "success", "message": msg, "data": note})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the note
	// note.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Artikel  string `json:"artikel"`
	}
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	// if note.ID == uuid.Nil {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	// }

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Artikel = updateNoteData.Artikel

	// Save the Changes
	db.Save(&note)

	// Return the updated note
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	// if note.ID == uuid.Nil {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	// }

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
