package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

const email = "fake"

type SpyStore struct {
	getCustomerEmailCalled bool
}

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
