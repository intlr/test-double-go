package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const email = "fake"

// MockStore describes a mocked datastore, a datastore which is doing some
// logic but also implements tools to validate expectations after the test
// execution.
type MockStore struct {
	methods map[string]bool
}

func NewStore() *MockStore {
	store := &MockStore{methods: make(map[string]bool)}

	return store
}

// ExpectToCall method adds a method to the list of expected methods to be
// called so that we can validate them after the test execution.
func (s *MockStore) ExpectToCall(method string) {
	s.methods[method] = false
}

// Verify method allows us to validate that methods were called.
func (s *MockStore) Verify(t *testing.T) {
	for method, called := range s.methods {
		if !called {
			t.Fatalf("method %q not called", method)
		}
	}
}

// GetCustomerEmail returns a customer email.
//
// The mocked function allows us to make sure the system under test would
// have called the production-ready function, as a spy would do. But we
// could go one step further and actually count the number of times the
// method was called, and expect the function to be called a very specific
// number of times. We could also implement some very specific scenarios as
// fake doubles.
func (s *MockStore) GetCustomerEmail(id int) string {
	s.methods["GetCustomerEmail"] = true

	return email
}

func TestService_Get(t *testing.T) {
	store := NewStore()
	serv := service.New(store)

	store.ExpectToCall("GetCustomerEmail")
	got := serv.Get()
	if got != email {
		t.Fatalf("got %q, want %q", got, email)
	}
	store.Verify(t)
}
