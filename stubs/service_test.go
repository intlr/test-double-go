package main

import "testing"

const email = "fake"

type StubStore struct {}

func (s StubStore) GetCustomerEmail(id int) string {
	return email
}

func TestService_Get(t *testing.T) {
	s := Service{store: StubStore{}}

	got := s.Get()
	if got != email {
		t.Fatalf("got %q, want %q", got, email)
	}
}
