package student

import (
	"database/sql"
	"testing"

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

}

func TestStudent_Get(t *testing.T) {

}

func TestStudent_Update(t *testing.T) {

}

func TestStudent_Delete(t *testing.T) {

}
