package main

import (
	"sync"
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const email = "fake"

var once sync.Once

type MockStore struct {
	methods map[string]bool
}

func NewStore() *MockStore {
	store := &MockStore{methods: make(map[string]bool)}

	return store
}

func (s *MockStore) ExpectToCall(method string) {
	s.methods[method] = false
}

func (s *MockStore) Verify(t *testing.T) {
	for method, called := range s.methods {
		if !called {
			t.Fatalf("method %q not called", method)
		}
	}
}

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
