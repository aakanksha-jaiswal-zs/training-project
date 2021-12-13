package employee

import "sample-api/services"

type handler struct {
	service services.Employee
}

/*
	Simple Factory:
	A factory function is a function that creates a new object, taking in the parameters required for its construction.
	It works much like a constructor. Factory functions are a better alternatives to initializing instances because,
	the function signature ensures that everyone will supply the required attributes.

	Factories are something that can be easily overused. If you are initializing a struct,
	where the default values in Go are also sensible defaults for you, then a factory isn’t really going to help.

	TIP: Use factory functions to ensure that new instances of structs are constructed with the required arguments
*/

// New is the factory function for handler that injects service dependency.
func New(service services.Employee) handler {
	return handler{service: service}
}
