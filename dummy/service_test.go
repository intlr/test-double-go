package main

import (
	"testing"

	"github.com/alr-lab/test-double/service"
)

type DummyStore struct{}

func (s DummyStore) GetCustomerEmail(id int) string {
	return ""
}

func TestServer_New(t *testing.T) {
	_ = service.New(DummyStore{})
}
