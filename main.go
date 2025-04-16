package main

import (
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen("0.0.0.0:" + port)
}
