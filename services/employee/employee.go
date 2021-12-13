package employee

import (
	"sample-api/filters"
	"sample-api/models"
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

/*
	The name of a method's receiver should be a reflection of its identity;
	often a one or two letter abbreviation of its type suffices (such as "c" or "cl" for "Client").
	Be consistent, too: if you call the receiver "c" in one method, don't call it "cl" in another.

	- Methods with value receivers can be called on pointers as well as values; Methods with pointer receivers
	 can only be called on pointers.
	- A value type creates a copy of the receiver when the method is invoked, so outside updates will
	 not be applied to this receiver. If the method needs to mutate the receiver, the receiver must be a pointer.
	- If the receiver is a large struct or array, a pointer receiver is more efficient.
	- Don't mix receiver types. Choose either pointers or struct types for all available methods.
*/

func (s service) Create(employee models.Employee) (models.Employee, error) {

}

func (s service) GetAll(filter filters.Employee) ([]models.Employee, error) {

}

func (s service) Get(id int64) (models.Employee, error) {

}

func (s service) Update(employee models.Employee) (models.Employee, error) {

}

func (s service) Delete(id int64) error {

}
