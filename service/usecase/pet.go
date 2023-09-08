package usecase

import (
	"context"

	"github.com/adrian-lin-1-0-0/go-ca/domain"
)

type PetUsecaseImpl struct {
	petRepo domain.PetRepository
}

type PetUsecaseOptions struct {
	PetRepo domain.PetRepository
}

func NewPetUsecase(opts *PetUsecaseOptions) domain.PetUsecase {
	return &PetUsecaseImpl{
		petRepo: opts.PetRepo,
	}
}

func (u *PetUsecaseImpl) GetPetById(ctx context.Context, id string) (*domain.Pet, error) {
	return nil, nil
}

func (u *PetUsecaseImpl) GetPet(ctx context.Context) ([]*domain.Pet, error) {
	return nil, nil
}

func (u *PetUsecaseImpl) AddPet(ctx context.Context, pet *domain.Pet) error {
	return nil
}
