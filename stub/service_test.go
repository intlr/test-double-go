package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const email = "fake"

// StubStore describes a stub datastore, a datastore which is very simple,
// without much logic.
type StubStore struct{}

// GetCustomerEmail returns a customer email.
//
// A stub method would simply return a fixed value in most cases.
func (s StubStore) GetCustomerEmail(id int) string {
	return email
}

func TestService_Get(t *testing.T) {
	serv := service.New(StubStore{})

	got := serv.Get()
	if got != email {
		t.Fatalf("got %q, want %q", got, email)
	}
}
