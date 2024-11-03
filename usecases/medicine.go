package usecases

import (
	"context"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"github.com/abrahammegantoro/quickcare-bpjs-be/repositories"
)

type MedicineUsecase struct {
	repo repositories.MedicineRepository
}

func NewMedicineUsecase(repo repositories.MedicineRepository) *MedicineUsecase {
	return &MedicineUsecase{
		repo: repo,
	}
}

func (u *MedicineUsecase) CreateMedicine(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error) {
	return u.repo.CreateMedicine(ctx, medicine)
}