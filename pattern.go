// Copyright (c) 2017 Thibault NORMAND
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package specifications

// Specification contract
type Specification interface {
	IsSatisfiedBy(object interface{}) bool
}

// -----------------------------------------------------------------------------

// AbstractSpecification implements Specification interface
type AbstractSpecification struct {
	Specification
}

// And returns a specification composition with AND operator
func (c *AbstractSpecification) And(other Specification) Specification {
	return &andSpecification{c, other}
}

// Or returns a specification composition with OR operator
func (c *AbstractSpecification) Or(other Specification) Specification {
	return &orSpecification{c, other}
}

// Not returns a specification composition with NOT operator
func (c *AbstractSpecification) Not() Specification {
	return &notSpecification{c}
}

// -----------------------------------------------------------------------------

type andSpecification struct {
	one Specification
	two Specification
}

func (a *andSpecification) IsSatisfiedBy(object interface{}) bool {
	return a.one.IsSatisfiedBy(object) && a.two.IsSatisfiedBy(object)
}

// -----------------------------------------------------------------------------

type orSpecification struct {
	one Specification
	two Specification
}

func (a *orSpecification) IsSatisfiedBy(object interface{}) bool {
	return a.one.IsSatisfiedBy(object) || a.two.IsSatisfiedBy(object)
}

// -----------------------------------------------------------------------------

type notSpecification struct {
	one Specification
}

func (a *notSpecification) IsSatisfiedBy(object interface{}) bool {
	return !a.one.IsSatisfiedBy(object)
}
