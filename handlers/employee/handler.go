package employee

import (
	"net/http"

	"sample-api/services"
)

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

/*
	Signature for Handler: func (w http.ResponseWriter, r *http.Request)

	https://pkg.go.dev/github.com/gorilla/mux#Router.HandleFunc
	The router's HandleFunc method takes a path [string] and a handler [func (w http.ResponseWriter, r *http.Request)]
	and routes requests to the respective handlers based on matching path.
*/

func (h handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {

}
