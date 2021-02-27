package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const email = "fake"

// SpyStore describes a spying datastore, a datastore which allows us to
// know if specific methods were called.
type SpyStore struct {
	getCustomerEmailCalled bool
}

// GetCustomerEmail returns a customer email.
//
// The spying method is recording that the production-ready method would
// have been called. An alternative is to actually count the number of
// times a method was called and expect it to be called a specific number
// of times after the test execution.
func (s *SpyStore) GetCustomerEmail(id int) string {
	s.getCustomerEmailCalled = true
	return email
}

func TestService_Get(t *testing.T) {
	spy := SpyStore{}
	serv := service.New(&spy)

	got := serv.Get()
	if got != email {
		t.Fatalf("got %q, want %q", got, email)
	}
	if !spy.getCustomerEmailCalled {
		t.Fatal("GetCustomerEmail not called")
	}
}
