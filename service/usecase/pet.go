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
	pet, err := u.petRepo.GetPetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func (u *PetUsecaseImpl) GetPet(ctx context.Context) ([]*domain.Pet, error) {
	pets, err := u.petRepo.GetPet(ctx)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (u *PetUsecaseImpl) AddPet(ctx context.Context, pet *domain.Pet) error {
	err := u.petRepo.AddPet(ctx, pet)
	if err != nil {
		return err
	}
	return nil
}

func (u *PetUsecaseImpl) DeletePet(ctx context.Context, id string) error {
	err := u.petRepo.DeletePet(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *PetUsecaseImpl) UpdatePet(ctx context.Context, pet *domain.Pet) error {
	err := u.petRepo.UpdatePet(ctx, pet)
	if err != nil {
		return err
	}
	return nil
}
