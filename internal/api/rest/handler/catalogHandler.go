package rest

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"net/http"
	"strconv"

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
	app.Get("/products", catalogHandler.GetAProduct)
	app.Get("/products/:id", catalogHandler.GetAllProducts)
	app.Get("/categories", catalogHandler.GetAllCategories)
	app.Get("/categories/:id", catalogHandler.GetACategory)

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
	selRoutes.Put("/products/:id", catalogHandler.StockUpdate) //update stock
	selRoutes.Delete("/products/:id", catalogHandler.DeleteProduct)

}

func (h *CatalogHandler) GetAllCategories(ctx *fiber.Ctx) error {

	allcat, err := h.svc.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err)
	}

	return rest.SuccessResponse(ctx, "All Category list-", allcat)
}

func (h *CatalogHandler) GetACategory(ctx *fiber.Ctx) error {
	//Extract id from URL
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "Invalid id parameter")
	}

	cat, err := h.svc.GetCategoryById(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err)
	}

	return rest.SuccessResponse(ctx, "Category for provided id", cat)
}

func (h *CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "create category invalid request")
	}

	cat, err := h.svc.CreateCategories(&req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Product category creation successfully completed", cat)
}

func (h *CatalogHandler) UpdateCategories(ctx *fiber.Ctx) error {
	//Extract id from URL
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "Invalid id parameter")
	}

	//Parse the body
	req := dto.CreateCategoryRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "invalid request for updating categories")
	}

	updatedCat, err := h.svc.UpdateCategories(id, &req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Category updation successfully completed", updatedCat)
}

func (h *CatalogHandler) DeleteCategories(ctx *fiber.Ctx) error {
	//Extract id from URL
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "Invalid id parameter")
	}

	if err := h.svc.DeleteCategories(id); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "category successfully deleted ", nil)
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
