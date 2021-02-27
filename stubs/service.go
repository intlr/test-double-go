package main

type (
	Service struct {
		store Store
	}

	Store interface {
		GetCustomerEmail(id int) string
	}
)

func (s Service) Get() string {
	return s.store.GetCustomerEmail(42)
}
