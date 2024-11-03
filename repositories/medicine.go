package repositories

import (
	"context"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MedicineRepository struct {
	collection *mongo.Collection
}

func NewMedicineRepository(db *mongo.Database) *MedicineRepository {
	return &MedicineRepository{
		collection: db.Collection("medicines"),
	}
}

func (r *MedicineRepository) CreateMedicine(ctx context.Context, medicine *models.Medicine) (*models.Medicine, error) {
	_, err := r.collection.InsertOne(ctx, medicine)
	if err != nil {
		return nil, err
	}

	return medicine, nil
}

func (r *MedicineRepository) GetAllMedicines(ctx context.Context) ([]*models.Medicine, error) {
	cursor, err := r.collection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	var medicines []*models.Medicine
	if err := cursor.All(ctx, &medicines); err != nil {
		return nil, err
	}

	return medicines, nil
}

func (r *MedicineRepository) GetMedicineByID(ctx context.Context, id string) (*models.Medicine, error) {
	medicine := &models.Medicine{}
	err := r.collection.FindOne(ctx, map[string]string{"_id": id}).Decode(medicine)
	if err != nil {
		return nil, err
	}

	return medicine, nil
}

func (r *MedicineRepository) UpdateMedicine(ctx context.Context, id string, update map[string]interface{}) error {
	_, err := r.collection.UpdateOne(ctx, map[string]string{"_id": id}, update)
	if err != nil {
		return err
	}

	return nil
}
