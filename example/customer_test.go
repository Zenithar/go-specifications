package example

import "testing"

func TestCustomerValidSpecification(t *testing.T) {
	customer := Customer{"Test 1", true}

	s := IsCustomerActive{}
	if !s.IsSatisfiedBy(customer) {
		t.Fail()
	}
}

func TestCustomerInvalidSpecification(t *testing.T) {
	customer := Customer{"Test 1", false}

	s := IsCustomerActive{}
	if s.IsSatisfiedBy(customer) {
		t.Fail()
	}
}
