package usecases

import (
	"context"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
)

type ReferralUsecases interface {
	CreateReferral(ctx context.Context, referral *models.Referral) (*models.Referral, error)
	GetAllReferrals(ctx context.Context) ([]*models.Referral, error)
	GetReferralByID(ctx context.Context, id string) (*models.Referral, error)
	UpdateReferral(ctx context.Context, id string, update map[string]interface{}) error
}

type ReferralUsecase struct {
	referralRepo ReferralUsecases
}

func NewReferralUsecase(referralRepo ReferralUsecases) *ReferralUsecase {
	return &ReferralUsecase{
		referralRepo: referralRepo,
	}
}

func (u *ReferralUsecase) CreateReferral(ctx context.Context, referral *models.Referral) (*models.Referral, error) {
	return u.referralRepo.CreateReferral(ctx, referral)
}

func (u *ReferralUsecase) GetAllReferrals(ctx context.Context) ([]*models.Referral, error) {
	return u.referralRepo.GetAllReferrals(ctx)
}

func (u *ReferralUsecase) GetReferralByID(ctx context.Context, id string) (*models.Referral, error) {
	return u.referralRepo.GetReferralByID(ctx, id)
}

func (u *ReferralUsecase) UpdateReferral(ctx context.Context, id string, update map[string]interface{}) error {
	return u.referralRepo.UpdateReferral(ctx, id, update)
}