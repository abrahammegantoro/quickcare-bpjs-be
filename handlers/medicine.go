package handlers

import (
	"context"
	"net/http"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"github.com/abrahammegantoro/quickcare-bpjs-be/usecases"
	"github.com/gofiber/fiber/v2"
)

type MedicineHandler struct {
	service usecases.MedicineUsecase
}

func NewMedicineHandler(service usecases.MedicineUsecase) *MedicineHandler {
	return &MedicineHandler{
		service: service,
	}
}	

func (h *MedicineHandler) CreateMedicine(c *fiber.Ctx) error {
	medicine := new(models.Medicine)
	if err := c.BodyParser(medicine); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := context.Background()
	medicine, err := h.service.CreateMedicine(ctx, medicine)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(medicine)
}