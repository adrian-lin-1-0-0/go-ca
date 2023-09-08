package domain

import "context"

type Pet struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

type PetUsecase interface {
	GetPetById(ctx context.Context, id string) (*Pet, error)
	GetPet(ctx context.Context) ([]*Pet, error)
	AddPet(ctx context.Context, pet *Pet) error
}

type PetRepository interface {
	GetPetById(ctx context.Context, id string) (*Pet, error)
	GetPet(ctx context.Context) ([]*Pet, error)
	AddPet(ctx context.Context, pet *Pet) error
}
