package routes

import (
	"github.com/ZishunD/moodgo/database"
	"github.com/ZishunD/moodgo/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/emotions", createEmotion)
	api.Get("/emotions", getAllEmotions)
	api.Delete("/emotions/:id", deleteEmotion)
}

// POST: 创建一条记录
func createEmotion(c *fiber.Ctx) error {
	emotion := new(models.Emotion)

	if err := c.BodyParser(emotion); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if emotion.Mood == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Mood is required",
		})
	}

	database.DB.Create(&emotion)
	return c.JSON(emotion)
}

// GET: get all emotions
func getAllEmotions(c *fiber.Ctx) error {
	var emotions []models.Emotion

	result := database.DB.Order("created_at desc").Find(&emotions)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve emotions",
		})
	}

	return c.JSON(emotions)
}

func deleteEmotion(c *fiber.Ctx) error {
	id := c.Params("id")

	var emotion models.Emotion

	result := database.DB.First(&emotion, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Emotion not found",
		})
	}

	database.DB.Delete(&emotion)

	return c.JSON(fiber.Map{
		"message": "Emotion deleted successfully",
	})
}
