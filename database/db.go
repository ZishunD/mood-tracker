package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// 载入 .env 文件
	if os.Getenv("GO_ENV") != "production" {
		_ = godotenv.Load()
	}

	// 从环境变量读取
	dsn := os.Getenv("DB_URL")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to database!")
	}

	DB = database
}
