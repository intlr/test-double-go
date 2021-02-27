package main

type (
	Service struct {
		store Store
	}

	Store interface {
		GetCustomerEmail(id int) string
	}
)

func New(store Store) *Service {
	return &Service{store: store}
}
