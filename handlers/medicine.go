package handlers

import (
	"context"
	"net/http"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"github.com/gofiber/fiber/v2"
)

type MedicineUsecase interface {
	CreateMedicine(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error)
	GetAllMedicines(ctx context.Context) ([]*models.Medicine, error)
	GetMedicineByID(ctx context.Context, id string) (*models.Medicine, error)
	UpdateMedicine(ctx context.Context, id string, update map[string]interface{}) error
}

type MedicineHandler struct {
	usecase MedicineUsecase
}

func NewMedicineHandler(fiber *fiber.Group, usecase MedicineUsecase) {
	handler := &MedicineHandler{
		usecase: usecase,
	}

	fiber.Post("/medicines", handler.CreateMedicine)
	fiber.Get("/medicines", handler.GetAllMedicines)
	fiber.Get("/medicines/:id", handler.GetMedicineByID)
	fiber.Put("/medicines/:id", handler.UpdateMedicine)
}

func (h *MedicineHandler) CreateMedicine(c *fiber.Ctx) error {
	medicine := new(models.Medicine)
	if err := c.BodyParser(medicine); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	medicine, err := h.usecase.CreateMedicine(ctx, medicine)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(medicine)
}

func (h *MedicineHandler) GetAllMedicines(c *fiber.Ctx) error {
	ctx := context.Background()
	medicines, err := h.usecase.GetAllMedicines(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(medicines)
}

func (h *MedicineHandler) GetMedicineByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	medicine, err := h.usecase.GetMedicineByID(ctx, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(medicine)
}

func (h *MedicineHandler) UpdateMedicine(c *fiber.Ctx) error {
	id := c.Params("id")
	update := make(map[string]interface{})
	if err := c.BodyParser(&update); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	err := h.usecase.UpdateMedicine(ctx, id, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
