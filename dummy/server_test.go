package main

import "testing"

type DummyDatastore struct{}

func (s DummyDatastore) GetCustomerNameByID(id int) string {
	return ""
}

func TestServer_New(t *testing.T) {
	_ = New(DummyDatastore{})
}
