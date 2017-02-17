package example

import (
	"zenithar.org/go/specifications"
)

// IsCustomerActive checks if customer is active
type IsCustomerActive struct {
	specifications.AbstractSpecification
}

// IsSatisfiedBy returns specification satisfaction status
func (s *IsCustomerActive) IsSatisfiedBy(object interface{}) bool {
	if customer, ok := object.(Customer); ok {
		return customer.Active
	}
	return false
}
