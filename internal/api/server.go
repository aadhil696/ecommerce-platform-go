package api

import (
	"go-ecommerce-app/configs"
	rest "go-ecommerce-app/internal/api/rest/handler"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection failed %v", err)
	}

	log.Println("Database connected successfully")

	//Run AutoMigration
	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	// log.Printf("Config DSN %v", config.Dsn)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
	}

	SetupRoutes(rh)

	app.Listen(config.ServerPort)
}

func SetupRoutes(rh *rest.RestHandler) {
	//user handler
	rest.SetupUserRoutes(rh)
	//transactions
	//catalog

}
