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
	app.Get("/products", catalogHandler.GetAllProducts)
	app.Get("/products/:id", catalogHandler.GetAProduct)
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
		return rest.BadRequestError(ctx, "Invalid id parameter", err)
	}

	cat, err := h.svc.GetCategoryById(id)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusInternalServerError, err)
	}

	return rest.SuccessResponse(ctx, "Category for provided id", cat)
}

func (h *CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	req := &dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(req); err != nil {
		return rest.BadRequestError(ctx, "create category invalid request", err)
	}

	cat, err := h.svc.CreateCategories(req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Product category creation successfully completed", cat)
}

func (h *CatalogHandler) UpdateCategories(ctx *fiber.Ctx) error {
	//Extract id from URL
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "Invalid id parameter", err)
	}

	//Parse the body
	req := &dto.CreateCategoryRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return rest.BadRequestError(ctx, "invalid request for updating categories", err)
	}

	updatedCat, err := h.svc.UpdateCategories(id, req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Category updation successfully completed", updatedCat)
}

func (h *CatalogHandler) DeleteCategories(ctx *fiber.Ctx) error {
	//Extract id from URL
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "Invalid id parameter", err)
	}

	if err := h.svc.DeleteCategories(id); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "category successfully deleted ", nil)
}

// Product handler Implementation
func (h *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	req := &dto.CreateProductRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return rest.BadRequestError(ctx, "Invalid request parameters", err)
	}

	//Getting current user for userid
	user := h.svc.Auth.GetCurrentUser(ctx) 
	prdct, err := h.svc.CreateProduct(user.ID, req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product created successfully", prdct)
}

func (h *CatalogHandler) GetAllProducts(ctx *fiber.Ctx) error {

	allprdcts, err := h.svc.GetAllProducts()
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "products list", allprdcts)
}

func (h *CatalogHandler) GetAProduct(ctx *fiber.Ctx) error {
	//Extract id from URL params
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "invalid parameter", err)
	}

	//Call service to get product using id
	prdct, err := h.svc.FindProductById(id)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product based on id", prdct)
}

func (h *CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	//Extract Product id from URL params
	prdctId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "invalid product id", err)
	}

	//Getting current userid for verifying product belongs to current seller
	user := h.svc.Auth.GetCurrentUser(ctx)

	req := &dto.CreateProductRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return rest.BadRequestError(ctx, "invalid request body params", err)
	}

	updatedPrdct, err := h.svc.UpdateProduct(prdctId, req, &user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product updated", updatedPrdct)
}

func (h *CatalogHandler) StockUpdate(ctx *fiber.Ctx) error {
	//Extract Product id from URL params
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "invalid product id", err)
	}

	req := &dto.UpdateStockRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return rest.BadRequestError(ctx, "invalid update request", err)
	}

	//Getting current userid for verifying product belongs to current seller
	user := h.svc.Auth.GetCurrentUser(ctx)
	updated, err := h.svc.StockUpdate(id, req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "stock updated successfully", updated)
}

func (h *CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	//Extract Product id from URL params
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "invalid product id", err)
	}

	if err := h.svc.DeleteProduct(id); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product deletion success", nil)
}
