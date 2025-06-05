package rest

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *RestHandler) {

	app := rh.App

	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	userHandler := UserHandler{
		svc: svc,
	}
	pubRoutes := app.Group("/users")
	//Public endpoints
	pubRoutes.Post("/register", userHandler.Register)
	pubRoutes.Post("/login", userHandler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	//Private endpoints
	pvtRoutes.Post("/verify", userHandler.Verify)
	pvtRoutes.Get("/verifycode", userHandler.GetVerificationCode)

	pvtRoutes.Post("/profile", userHandler.CreateProfile)
	pvtRoutes.Get("/profile", userHandler.GetProfile)

	pvtRoutes.Post("/cart", userHandler.AddToCart)
	pvtRoutes.Get("/cart", userHandler.GetCart)
	pvtRoutes.Get("/order", userHandler.Getorders)
	pvtRoutes.Get("/order/:id", userHandler.GetOrder)

	pvtRoutes.Post("/become-seller", userHandler.BecomeSeller)

}

func (h UserHandler) Register(ctx *fiber.Ctx) error {
	//to create user
	user := &dto.UserSignUp{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
		})
	}

	token, err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Internal error on signup",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Registration successfull",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	user := &dto.UserLogin{}
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid Login details",
		})
	}

	token, err := h.svc.Login(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Login failed due to some internal error",
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	//Current User
	user := h.svc.Auth.GetCurrentUser(ctx)

	var req *dto.UserVerifyCode

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Json parsing error",
		})
	}

	if err := h.svc.VerifyCode(user.ID, req.Code); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Verification successfull",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	code, err := h.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "code generating failed",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Verified",
		"data":    code,
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	//Get current users details for profile creation
	user := h.svc.Auth.GetCurrentUser(ctx)

	req := &dto.ProfileInput{}
	if err := ctx.BodyParser(req); err != nil {
		return rest.BadRequestError(ctx, "invalid input parameters", err)
	}

	if err := h.svc.CreateProfile(uint(user.ID), req); err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Profile created",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)

	//call service layer method to get current user profile
	profile, err := h.svc.GetProfile(uint(user.ID))
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":     "Profile",
		"userprofile": profile,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Added to cart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "GotCart",
	})
}

func (h *UserHandler) Getorders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Got orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get order",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request parameters",
		})
	}

	token, err := h.svc.BecomeSeller(user.ID, req)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "failed to upgrade as seller",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "seller mode accomplished",
		"token":   token,
	})
}
