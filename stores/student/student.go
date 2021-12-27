package student

import (
	"database/sql"

	"sample-api/filters"
	"sample-api/models"
	"sample-api/stores"
)

type store struct {
	db *sql.DB
}

// New is the factory function for stores that injects database dependency.
func New(db *sql.DB) stores.Student {
	return store{db: db}
}

func (s store) Create(student models.Student) error {
	return nil
}

func (s store) GetAll(filter filters.Student) ([]models.Student, error) {
	return nil, nil
}

func (s store) Get(id int64) (models.Student, error) {
	return models.Student{}, nil
}

func (s store) Update(student models.Student) error {
	return nil
}

func (s store) Delete(id int64) error {
	return nil
}
