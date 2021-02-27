package main

import (
	"testing"

	"github.com/alr-lab/test-double/service"
)

const email = "fake"

type StubStore struct {}

func (s StubStore) GetCustomerEmail(id int) string {
	return email
}

func TestService_Get(t *testing.T) {
	s := service.New(StubStore{})

	got := s.Get()
	if got != email {
		t.Fatalf("got %q, want %q", got, email)
	}
}
