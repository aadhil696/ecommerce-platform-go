package rest

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *RestHandler) {

	app := rh.App

	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	catalogHandler := CatalogHandler{
		svc: svc,
	}

	//public
	//listing products and categories
	app.Get("/products")
	app.Get("products/:id")
	app.Get("/categories")
	app.Get("/categories/:id")

	//private
	//manage products and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)

	//categories
	selRoutes.Post("/categories", catalogHandler.CreateCategories)
	selRoutes.Patch("/categories/:id", catalogHandler.UpdateCategories)
	selRoutes.Delete("/categories/:id", catalogHandler.DeleteCategories)

	//products
	selRoutes.Post("/products", catalogHandler.CreateProduct)
	selRoutes.Get("/products", catalogHandler.GetAllProducts)
	selRoutes.Get("/products/:id", catalogHandler.GetAProduct)
	selRoutes.Patch("/products/:id", catalogHandler.EditProduct)
	selRoutes.Put("/product/:id", catalogHandler.StockUpdate) //update stock
	selRoutes.Delete("/products/:id", catalogHandler.DeleteProduct)

}

func (h *CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "create category invalid request")
	}

	if err := h.svc.CreateCategories(&req); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "PRoduct category successfully completed", nil)
}

func (h *CatalogHandler) UpdateCategories(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "invalid request for updating categories")
	}

	if err := h.svc.UpdateCategories(&req); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Category updation successfully completed", nil)
}

func (h *CatalogHandler) DeleteCategories(ctx *fiber.Ctx) error {

	
	return rest.SuccessResponse(ctx, "delete catalog ", nil)
}

func (h *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "product created", nil)
}

func (h *CatalogHandler) GetAllProducts(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "product view", nil)
}

func (h *CatalogHandler) GetAProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "product based on id", nil)
}

func (h *CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "product updation", nil)
}

func (h *CatalogHandler) StockUpdate(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "stock updation", nil)
}

func (h *CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "product delete", nil)
}
