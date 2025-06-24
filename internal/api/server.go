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
	err = db.AutoMigrate(&domain.User{},
		&domain.BankAccount{},
		&domain.Category{},
		&domain.Product{},
		&domain.Cart{},
		&domain.Address{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Payment{},
	)
	if err != nil {
		log.Fatalf("Migration failed due to %s", err)
	}

	log.Println("Migration was successfull")

	// //cors configuration
	// c := cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:3000",
	// 	AllowHeaders: "Content-Type, Accept, Authorization",
	// 	AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	// })

	// app.Use(c)

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
	rest.SetupCatalogRoutes(rh)

}
