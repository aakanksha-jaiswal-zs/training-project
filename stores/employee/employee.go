package employee

import (
	"database/sql"

	"sample-api/filters"
	"sample-api/models"
	"sample-api/stores"
)

/*
	As a practice, we group imports from same domain. This makes the imports look clean.
	All the files should follow the same pattern for imports ordering. Each group should be in alphabetical order.

	Order for grouping imports:
	1. Standard libraries
	2. Groups of external libraries
	3. Local imports
*/

type store struct {
	db *sql.DB
}

/*
	Interface Factory:
	Factory functions can return interfaces instead of structs.
	Interfaces allow you to define behavior without exposing internal implementation.
	This means we can make structs private, while only exposing the interface outside our package.
*/

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
