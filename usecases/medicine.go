package usecases

import (
	"context"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
)

type MedicineRepository interface {
	CreateMedicine(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error)
	GetAllMedicines(ctx context.Context) ([]*models.Medicine, error)
	GetMedicineByID(ctx context.Context, id string) (*models.Medicine, error)
	UpdateMedicine(ctx context.Context, id string, update map[string]interface{}) error
}

type MedicineUsecase struct {
	medicineRepo MedicineRepository
}

func NewMedicineUsecase(medicineRepo MedicineRepository) *MedicineUsecase {
	return &MedicineUsecase{
		medicineRepo: medicineRepo,
	}
}

func (u *MedicineUsecase) CreateMedicine(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error) {
	return u.medicineRepo.CreateMedicine(ctx, medicine)
}

func (u *MedicineUsecase) GetAllMedicines(ctx context.Context) ([]*models.Medicine, error) {
	return u.medicineRepo.GetAllMedicines(ctx)
}

func (u *MedicineUsecase) GetMedicineByID(ctx context.Context, id string) (*models.Medicine, error) {
	return u.medicineRepo.GetMedicineByID(ctx, id)
}

func (u *MedicineUsecase) UpdateMedicine(ctx context.Context, id string, update map[string]interface{}) error {
	return u.medicineRepo.UpdateMedicine(ctx, id, update)
}