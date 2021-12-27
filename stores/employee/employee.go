package employee

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
func New(db *sql.DB) stores.Employee {
	return store{db: db}
}

func (s store) Create(employee models.Employee) (int64, error) {
	return 0, nil
}

func (s store) GetAll(filter filters.Employee) ([]models.Employee, error) {
	return nil, nil
}

func (s store) Get(id int64) (models.Employee, error) {
	return models.Employee{}, nil
}

func (s store) Update(employee models.Employee) error {
	return nil
}

func (s store) Delete(id int64) error {
	return nil
}
