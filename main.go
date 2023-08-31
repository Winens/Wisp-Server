package main

import (
	"github.com/Winens/Wisp-Server/db"
	"github.com/Winens/Wisp-Server/model"
	"github.com/Winens/Wisp-Server/pkg/session"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func connectDatabases() {
	db.ConnectPostgreSQL()
}

func migrateDatabases() {
	if err := db.PostgeSQL.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	_ = godotenv.Load(".env")

	app := fiber.New(fiber.Config{
		AppName: "Wisp - Wi Secure Protocol",

		// Faster JSON serialization
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		CaseSensitive: true,
	})

	connectDatabases()
	migrateDatabases()

	// Middlewares
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("SECRET_KEY"),
	}))

	session.Init()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
