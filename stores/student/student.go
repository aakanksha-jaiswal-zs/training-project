package student

import (
	"sample-api/stores"
)

type store struct {
}

/*
	Interface Factory:
	Factory functions can return interfaces instead of structs.
	Interfaces allow you to define behavior without exposing internal implementation.
	This means we can make structs private, while only exposing the interface outside our package.
*/

// New is the factory function for stores that injects database dependency.
func New() stores.Student {
	return store{}
}
