package repositories

import (
	"context"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReferralRepository struct {
	collection *mongo.Collection
}

func NewReferralRepository(db *mongo.Database) *ReferralRepository {
	return &ReferralRepository{
		collection: db.Collection("referrals"),
	}
}

func (r *ReferralRepository) CreateReferral(ctx context.Context, referral *models.Referral) (*models.Referral, error) {
	_, err := r.collection.InsertOne(ctx, referral)
	if err != nil {
		return nil, err
	}

	return referral, nil
}

func (r *ReferralRepository) GetAllReferrals(ctx context.Context) ([]*models.Referral, error) {
	cursor, err := r.collection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	var referrals []*models.Referral
	if err := cursor.All(ctx, &referrals); err != nil {
		return nil, err
	}

	return referrals, nil
}

func (r *ReferralRepository) GetReferralByID(ctx context.Context, id string) (*models.Referral, error) {
	referral := &models.Referral{}
	err := r.collection.FindOne(ctx, map[string]string{"_id": id}).Decode(referral)
	if err != nil {
		return nil, err
	}

	return referral, nil
}

func (r *ReferralRepository) UpdateReferral(ctx context.Context, id string, update map[string]interface{}) error {
	_, err := r.collection.UpdateOne(ctx, map[string]string{"_id": id}, update)
	if err != nil {
		return err
	}

	return nil
}
