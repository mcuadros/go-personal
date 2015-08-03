package personal

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type PersonalSuite struct{}

var _ = Suite(&PersonalSuite{})

var fullNameProvider = map[string][]string{
	"Máximo Cuadros Ortiz": []string{"mcuadros", "Máximo Cuadros", "Máximo Cuadros Ortiz"},
	"Foo Bar":              []string{"foo", "foo bar", "foo ba", "Foo Bar"},
	"Foo Bar Qux":          []string{"Foo Bar Qux", "Foo Bar"},
	"Bar Bar Qux":          []string{"Bar Bar Qux", "Bar Bar Qux Bar"},
	"Qux Bar":              []string{"foo bar Qux", "Qux Bar"},
	"Qux Bazzz":            []string{"Qux Bazzz", "Qux Bar"},
	"Bar Baz":              []string{"Bar Baz", "B2ar Baz"},
	"Foo":                  []string{"QUX", "Foo"},
	"Foo Qux":              []string{"Foo Qux", "Foo Qux (foo)"},
	"Bär Foo":              []string{"Bär Foo", "Bar Foo"},
	"Bár Foo":              []string{"Bár Foo", "B�r Foo"},
	"Bar Foo":              []string{"Bar Foo", "B�r Foo"},
	"Foo Garcia-Castro":    []string{"Foo", "Foo Garcia-Castro"},
}

func (s *PersonalSuite) TestGetBestFullName(c *C) {
	for best, names := range fullNameProvider {
		c.Assert(best, Equals, GetBestFullName(names))
	}
}
