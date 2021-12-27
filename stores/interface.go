package stores

import (
	"sample-api/filters"
	"sample-api/models"
)

type Employee interface {
	Create(employee models.Employee) (int64, error)
	GetAll(filter filters.Employee) ([]models.Employee, error)
	Get(id int64) (models.Employee, error)
	Update(employee models.Employee) error
	Delete(id int64) error
}

type Student interface {
	Create(student models.Student) error
	GetAll(filter filters.Student) ([]models.Student, error)
	Get(id int64) (models.Student, error)
	Update(student models.Student) error
	Delete(id int64) error
}
