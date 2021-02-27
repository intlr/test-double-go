package main

type (
	// Server describes any kind of server. Its purpose is to demonstrate
	// dummy doubles.
	Server struct {
		store Datastore
	}

	// Datastore describes a datastore contract
	Datastore interface {
		GetCustomerNameByID(id int) string
	}
)

// New returns a new server
func New(store Datastore) *Server {
	return &Server{store: store}
}
