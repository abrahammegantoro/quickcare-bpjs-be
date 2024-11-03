package handlers

import (
	"context"
	"net/http"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"github.com/gofiber/fiber/v2"
)

type ReferralUsecase interface {
	CreateReferral(ctx context.Context, referral *models.Referral) (*models.Referral, error)
	GetAllReferrals(ctx context.Context) ([]*models.Referral, error)
	GetReferralByID(ctx context.Context, id string) (*models.Referral, error)
	UpdateReferral(ctx context.Context, id string, update map[string]interface{}) error
}

type ReferralHandler struct {
	usecase ReferralUsecase
}

func NewReferralHandler(fiber *fiber.Group, usecase ReferralUsecase) {
	handler := &ReferralHandler{
		usecase: usecase,
	}

	fiber.Post("/referrals", handler.CreateReferral)
	fiber.Get("/referrals", handler.GetAllReferrals)
	fiber.Get("/referrals/:id", handler.GetReferralByID)
	fiber.Put("/referrals/:id", handler.UpdateReferral)
}

func (h *ReferralHandler) CreateReferral(c *fiber.Ctx) error {
	referral := new(models.Referral)
	if err := c.BodyParser(referral); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	referral, err := h.usecase.CreateReferral(ctx, referral)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(referral)
}

func (h *ReferralHandler) GetAllReferrals(c *fiber.Ctx) error {
	ctx := context.Background()
	referrals, err := h.usecase.GetAllReferrals(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(referrals)
}

func (h *ReferralHandler) GetReferralByID(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	referral, err := h.usecase.GetReferralByID(ctx, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(referral)
}

func (h *ReferralHandler) UpdateReferral(c *fiber.Ctx) error {
	id := c.Params("id")
	update := make(map[string]interface{})
	if err := c.BodyParser(&update); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	err := h.usecase.UpdateReferral(ctx, id, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
