package rest

import (
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func InitializeTransactionService(db *gorm.DB, auth helper.Auth) service.TransactionService {
	return service.TransactionService{
		Repo: repository.NewTransactionRepo(db),
		Auth: auth,
	}
}

func SetupTransactionRoutes(rh *RestHandler) {

	app := rh.App
	svc := InitializeTransactionService(rh.DB, rh.Auth)

	handler := TransactionHandler{
		svc: svc,
	}

	secRoute := app.Group("/", rh.Auth.Authorize)
	secRoute.Get("/payment", handler.MakePayment)

	sellerRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	sellerRoutes.Get("/orders", handler.GetOrders)
	sellerRoutes.Get("/orders/:id", handler.GetOrderDetails)

}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	return nil
}

func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	return nil
}
func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	return nil
}
