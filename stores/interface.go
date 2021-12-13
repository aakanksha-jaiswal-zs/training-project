package stores

import (
	"sample-api/filters"
	"sample-api/models"
)

type Employee interface {
	Create(employee models.Employee) (models.Employee, error)
	GetAll(filter filters.Employee) ([]models.Employee, error)
	Get(id int64) (models.Employee, error)
	Update(employee models.Employee) (models.Employee, error)
	Delete(id int64) error
}

type Student interface {
	Create(student models.Student) (models.Student, error)
	GetAll(filter filters.Student) ([]models.Student, error)
	Get(id int64) (models.Student, error)
	Update(student models.Student) (models.Student, error)
	Delete(id int64) error
}
