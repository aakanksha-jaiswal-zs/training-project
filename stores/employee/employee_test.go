package employee

import (
	"database/sql"
	"testing"

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
		id    int64 // output
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

}

func TestEmployee_Get(t *testing.T) {

}

func TestEmployee_Update(t *testing.T) {

}

func TestEmployee_Delete(t *testing.T) {

}
