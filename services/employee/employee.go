package employee

import (
	"sample-api/services"
	"sample-api/stores"
)

type service struct {
	store stores.Employee
}

/*
	Interface Factory:
	Factory functions can return interfaces instead of structs.
	Interfaces allow you to define behavior without exposing internal implementation.
	This means we can make structs private, while only exposing the interface outside our package.
*/

// New is the factory function for service that injects store dependency.
func New(store stores.Employee) services.Employee {
	return service{store: store}
}
