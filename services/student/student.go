package student

import (
	"sample-api/services"
	"sample-api/stores"
)

type service struct {
	store stores.Student
}

/*
	Interface Factory:
	Factory functions can return interfaces instead of structs.
	Interfaces allow you to define behavior without exposing internal implementation.
	This means we can make structs private, while only exposing the interface outside our package.
*/

// New is the factory function for service that injects store dependency.
func New(store stores.Student) services.Student {
	return service{store: store}
}
