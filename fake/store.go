package main

type Store interface {
	GetCustomerEmail(id int) string
}
