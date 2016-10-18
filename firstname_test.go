package personal

import (
	. "gopkg.in/check.v1"
)

var firstNameProvider = map[string][]string{
	"Máximo": {"mcuadros", "Máximo Cuadros", "Máximo Cuadros Ortiz"},
	"Foo":    {"foo", "foo bar", "foo ba", "Foo Bar"},
	"":       {"QUX", "Foo"},
	"Bar":    {"Bar Foo", "B�r Foo"},
	"Qux":    {"A. Qux Foo", "B�r Foo"},
	"Peter":  {"Mr. Peter Foo"},
	"Helen":  {"Mrs. Helen Foo"},
	"Who":    {"Dr. Who Foo"},
	"Whoa":   {"Prof Whoa Foo"},
}

func (s *PersonalSuite) TestGetBestFirstname(c *C) {
	for best, names := range firstNameProvider {
		c.Assert(GetBestFirstname(names), Equals, best)
	}
}
