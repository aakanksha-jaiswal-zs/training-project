package employee

import (
	"database/sql"
	"reflect"
	"testing"

	"sample-api/filters"
	"sample-api/models"
	"sample-api/stores"
)

func initializeTests() stores.Employee {
	var db *sql.DB

	return New(db)
}

func TestEmployee_Create(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc  string
		input models.Employee
		id    int64
	}{
		{"create with missing optional fields", models.Employee{Name: "Joey", Role: models.NonFaculty}, 1},
		{"create with all fields", models.Employee{Name: "Jane", Designation: "Admin", Role: models.NonFaculty}, 2},
		{"employee ID ignored for create request", models.Employee{ID: 10, Name: "Alex"}, 3},
	}

	for i, tc := range tests {
		id, err := s.Create(tc.input)

		if err != nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: nil, Got: %v", i, tc.desc, err)
		}

		if id != tc.id {
			t.Errorf("TEST[%d], failed: %s\nExpected: %v, Got: %v", i, tc.desc, tc.id, id)
		}
	}
}

func TestEmployee_CreateError(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc  string
		input models.Employee
	}{
		{"missing mandatory fields", models.Employee{Designation: "Admin"}},
		{"database error", models.Employee{Name: "Chandler", Designation: "Professor"}},
	}

	for i, tc := range tests {
		id, err := s.Create(tc.input)

		if err == nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: error, Got: nil", i, tc.desc)
		}

		if id != 0 {
			t.Errorf("TEST[%d], failed: %s\nExpected: no ID to be created, Got: %v", i, tc.desc, id)
		}
	}
}

func TestEmployee_GetAll(t *testing.T) {
	s := initializeTests()

	employees := []models.Employee{
		{ID: 10, Name: "Bryce", Designation: "Accountant", Role: models.NonFaculty},
		{ID: 20, Name: "Clay", Designation: "Registrar", Role: models.NonFaculty},
		{ID: 12, Name: "Hannah", Designation: "Assistant Professor"},
	}

	tests := []struct {
		desc   string
		filter filters.Employee
		output []models.Employee
	}{
		{"list all employees", filters.Employee{}, employees},
		{"list all non-faculty employees", filters.Employee{Role: models.NonFaculty}, employees[:2]},
	}

	for i, tc := range tests {
		output, err := s.GetAll(tc.filter)

		if err != nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: nil, Got: %v", i, tc.desc, err)
		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("TEST[%d], failed: %s\nExpected: %v,\nGot: %v", i, tc.desc, tc.output, output)
		}
	}
}

func TestEmployee_GetAllError(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc   string
		filter filters.Employee
	}{
		{"database error", filters.Employee{}},
	}

	for i, tc := range tests {
		output, err := s.GetAll(tc.filter)

		if err == nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: error, Got: nil", i, tc.desc)
		}

		if output != nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: nil,\nGot: %v", i, tc.desc, output)
		}
	}
}

func TestEmployee_Get(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc   string
		id     int64
		output *models.Employee
	}{
		{"get employee with ID 10", 10, &models.Employee{ID: 10, Name: "Bryce", Designation: "Accountant", Role: models.NonFaculty}},
	}

	for i, tc := range tests {
		output, err := s.Get(tc.id)

		if err != nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: nil, Got: %v", i, tc.desc, err)
		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("TEST[%d], failed: %s\nExpected: %v,\nGot: %v", i, tc.desc, tc.output, output)
		}
	}
}

func TestEmployee_GetError(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc string
		id   int64
	}{
		{"db error", 20},
		{"entity does not exist", 999},
	}

	for i, tc := range tests {
		output, err := s.Get(tc.id)

		if err == nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: error, Got: nil", i, tc.desc)
		}

		if !reflect.DeepEqual(models.Employee{}, output) {
			t.Errorf("TEST[%d], failed: %s\nExpected: %v ,\nGot: %v", i, tc.desc, models.Employee{}, output)
		}
	}
}

func TestEmployee_Update(t *testing.T) {

}

func TestEmployee_Delete(t *testing.T) {

}
