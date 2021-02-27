package main

type Service struct {
	store Store
}

func (s Service) Get() string {
	return s.store.GetCustomerEmail(42)
}
