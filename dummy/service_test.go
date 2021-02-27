package main

import "testing"

type DummyStore struct{}

func (s DummyStore) GetCustomerEmail(id int) string {
	return ""
}

func TestServer_New(t *testing.T) {
	_ = New(DummyStore{})
}
