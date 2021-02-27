package main

import (
	"testing"

	"github.com/alr-lab/test-double-go/service"
)

// DummyStore describes a dummy datastore, a datastore which is not even
// being used but passed as an argument to allow the test execution flow to
// proceed.
type DummyStore struct{}

// GetCustomerEmail returns a customer email.
//
// This implementation allows to pass the DummyStore as a storage to the
// service but it is not being called during the test.
func (s DummyStore) GetCustomerEmail(id int) string {
	return ""
}

func TestServer_New(t *testing.T) {
	_ = service.New(DummyStore{})
}
