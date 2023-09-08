package mysql

import (
	"context"
	"database/sql"

	"github.com/adrian-lin-1-0-0/go-ca/domain"
)

type Pet struct {
	db *sql.DB
}

func NewPetRepository(db *sql.DB) domain.PetRepository {
	return &Pet{
		db: db,
	}
}

func (p *Pet) GetPetById(ctx context.Context, id string) (*domain.Pet, error) {
	row := p.db.QueryRow("SELECT * FROM pet WHERE id = $1", id)
	var pet domain.Pet
	err := row.Scan(&pet.Id, &pet.Name)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (p *Pet) GetPet(ctx context.Context) ([]*domain.Pet, error) {
	rows, err := p.db.Query("SELECT * FROM pet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []*domain.Pet
	for rows.Next() {
		var pet domain.Pet
		err := rows.Scan(&pet.Id, &pet.Name)
		if err != nil {
			return nil, err
		}
		pets = append(pets, &pet)
	}
	return pets, nil
}

func (p *Pet) AddPet(ctx context.Context, pet *domain.Pet) error {
	_, err := p.db.Exec("INSERT INTO pet (name) VALUES ($1)", pet.Name)
	if err != nil {
		return err
	}
	return nil
}
