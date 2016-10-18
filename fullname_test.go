package personal

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type PersonalSuite struct{}

var _ = Suite(&PersonalSuite{})

var fullNameProvider = map[string][]string{
	"Máximo Cuadros Ortiz": {"mcuadros", "Máximo Cuadros", "Máximo Cuadros Ortiz"},
	"Foo Bar":              {"foo", "foo bar", "foo ba", "Foo Bar"},
	"Foo Bar Qux":          {"Foo Bar Qux", "Foo Bar"},
	"Bar Bar Qux Bar":      {"Bar Bar Qux", "Bar Bar Qux Bar"},
	"Qux Bar":              {"foo bar Qux", "Qux Bar"},
	"Qux Bazzz":            {"Qux Bazzz", "Qux Bar"},
	"Bar Baz":              {"Bar Baz", "B2ar Baz"},
	"Foo":                  {"QUX", "Foo"},
	"Foo Qux":              {"Foo Qux", "Foo Qux (foo)"},
	"Bär Foo":              {"Bär Foo", "Bar Foo"},
	"Bár Foo":              {"Bár Foo", "B�r Foo"},
	"Bar Foo":              {"Bar Foo", "B�r Foo"},
	"John J. Aa":           {"John Aa", "John J.", "John J. Aa"},
	"John J Ab":            {"John Ab", "John J.", "John J Ab"},
	"John J. Ac":           {"John J. Ac", "John J Ac"},
	"藤原藤原藤原":               {"藤原藤原藤原", "2005"},
	"Foo Garcia-Castro":    {"Foo", "Foo Garcia-Castro"},
	"Foo de la O":          {"Foo del Bar", "Foo de la O", "Foo Bar", "Foo de la O blah"},
	"Foo Bar Baz":          {"Fo Ba Ba", "Foo Bar Baz"},
	"Foo Bar Bar":          {"Foo Bar Bar", "Foo Bar Baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaar"},
	"bpauls":               {"bpauls", "2012-01-13"},
	"Peter Bar III":        {"Peter", "Peter Bar", "Peter Bar III"},
	"Peter III":            {"Peter", "Peter III", "Peter III Bar"},
}

func (s *PersonalSuite) TestGetBestFullName(c *C) {
	for best, names := range fullNameProvider {
		c.Assert(GetBestFullName(names), Equals, best)
	}
}
