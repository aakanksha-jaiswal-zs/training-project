package student

import (
	"database/sql"
	"reflect"
	"testing"

	"sample-api/filters"
	"sample-api/models"
	"sample-api/stores"
)

func initializeTests() stores.Student {
	var db *sql.DB // fixme: initialize mock database

	return New(db)
}

func TestStudent_Create(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc    string
		student models.Student
	}{
		{"missing optional fields", models.Student{EnrollmentNumber: 1, Name: "James", CurrentSemester: 6}},
		{"create with all fields", models.Student{EnrollmentNumber: 2, Name: "Phoebe", CurrentSemester: 5, CGPA: 8.0, SpeciallyAble: false}},
	}

	for i, tc := range tests {
		err := s.Create(tc.student)

		if err != nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: nil, Got: %v", i, tc.desc, err)
		}
	}
}

func TestStudent_CreateError(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc    string
		student models.Student
	}{
		{"missing primary key", models.Student{Name: "Alex"}},
		{"missing mandatory field (Name)", models.Student{EnrollmentNumber: 1, CurrentSemester: 6}},
		{"database error", models.Student{EnrollmentNumber: 102, Name: "Alex", CurrentSemester: 5}},
	}

	for i, tc := range tests {
		err := s.Create(tc.student)

		if err == nil {
			t.Errorf("TEST[%d], failed: %s\nExpected: error, Got: nil", i, tc.desc)
		}
	}
}

func TestStudent_GetAll(t *testing.T) {
	s := initializeTests()

	students := []models.Student{
		{EnrollmentNumber: 10, Name: "Bryce", CurrentSemester: 2, CGPA: 5.8},
		{EnrollmentNumber: 20, Name: "Clay", CurrentSemester: 2, CGPA: 8.8},
		{EnrollmentNumber: 12, Name: "Hannah", CurrentSemester: 4, CGPA: 9.0},
	}

	tests := []struct {
		desc   string
		filter filters.Student
		output []models.Student
	}{
		{"list all students", filters.Student{}, students},
		{"list students of first year", filters.Student{Year: 1}, students[:2]},
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

func TestStudent_GetAllError(t *testing.T) {
	s := initializeTests()

	tests := []struct {
		desc   string
		filter filters.Student
	}{
		{"database error", filters.Student{}},
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

func TestStudent_Get(t *testing.T) {

}

func TestStudent_Update(t *testing.T) {

}

func TestStudent_Delete(t *testing.T) {

}
