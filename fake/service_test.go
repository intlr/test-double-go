package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const (
	emailDefault = ""
	emailValidUser = "fake"
)

type FakeStore struct {}

func (s FakeStore) GetCustomerEmail(id int) string {
	email := emailDefault
	switch id {
	case 42:
		email = emailValidUser
	}

	return email
}

func TestService_Get(t *testing.T) {
	s := service.New(FakeStore{})

	got := s.Get()
	if got != emailValidUser {
		t.Fatalf("got %q, want %q", got, emailValidUser)
	}
}
