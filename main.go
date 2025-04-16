package main

import (
	"github.com/ZishunD/moodgo/database"
	"github.com/ZishunD/moodgo/models"
	"github.com/ZishunD/moodgo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	// 自动建表
	database.DB.AutoMigrate(&models.Emotion{})

	// 注册路由
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
