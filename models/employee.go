package models

type Employee struct {
	ID          int64
	Name        string
	Designation string
	Role        Role
}

type Role int

const (
	Faculty Role = iota
	NonFaculty
)
