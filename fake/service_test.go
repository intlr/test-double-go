package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const (
	emailDefault = ""
	emailValidUser = "fake"
)

// FakeStore describes a fake datastore, a datastore which is doing some
// kind of logic but not the one that we might expect in production. It
// allows us to take shortcuts and return errors to validate the behaviour
// of the system under test during the test execution.
type FakeStore struct {}

// GetCustomerEmail returns a customer email.
//
// Here there is a simple logic of testing its argument and returning data
// accordingly. We could also return an error or add more data to be
// returned.
func (s FakeStore) GetCustomerEmail(id int) string {
	email := emailDefault
	switch id {
	case 42:
		email = emailValidUser
	}

	return email
}

func TestService_Get(t *testing.T) {
	serv := service.New(FakeStore{})

	got := serv.Get()
	if got != emailValidUser {
		t.Fatalf("got %q, want %q", got, emailValidUser)
	}
}
