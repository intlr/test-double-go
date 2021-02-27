package main

import "testing"

const expectedEmail = "fake"

type FakeStore struct {}

func (s FakeStore) GetCustomerEmail(id int) string {
	return expectedEmail
}

func TestService_Get(t *testing.T) {
	s := Service{store: FakeStore{}}

	got := s.Get()
	if got != expectedEmail {
		t.Fatalf("got %q, want %q", got, expectedEmail)
	}
}
