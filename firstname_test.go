package personal

import (
	. "gopkg.in/check.v1"
)

var firstNameProvider = map[string][]string{
	"Máximo": []string{"mcuadros", "Máximo Cuadros", "Máximo Cuadros Ortiz"},
	"Foo":    []string{"foo", "foo bar", "foo ba", "Foo Bar"},
	"":       []string{"QUX", "Foo"},
	"Bar":    []string{"Bar Foo", "B�r Foo"},
	"Qux":    []string{"A. Qux Foo", "B�r Foo"},
}

func (s *PersonalSuite) TestGetBestFirstname(c *C) {
	for best, names := range firstNameProvider {
		c.Assert(best, Equals, GetBestFirstname(names))
	}
}
