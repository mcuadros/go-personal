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
	"Bar Bar Qux Bar":      []string{"Bar Bar Qux", "Bar Bar Qux Bar"},
	"Qux Bar":              []string{"foo bar Qux", "Qux Bar"},
	"Qux Bazzz":            []string{"Qux Bazzz", "Qux Bar"},
	"Bar Baz":              []string{"Bar Baz", "B2ar Baz"},
	"Foo":                  []string{"QUX", "Foo"},
	"Foo Qux":              []string{"Foo Qux", "Foo Qux (foo)"},
	"Bär Foo":              []string{"Bär Foo", "Bar Foo"},
	"Bár Foo":              []string{"Bár Foo", "B�r Foo"},
	"Bar Foo":              []string{"Bar Foo", "B�r Foo"},
	"John J. Aa":           []string{"John Aa", "John J.", "John J. Aa"},
	"John J Ab":            []string{"John Ab", "John J.", "John J Ab"},
	"John J. Ac":           []string{"John J. Ac", "John J Ac"},
	"藤原藤原藤原":         []string{"藤原藤原藤原", "2005"},
	"Foo Garcia-Castro":    []string{"Foo", "Foo Garcia-Castro"},
	"Foo de la O":          []string{"Foo del Bar", "Foo de la O", "Foo Bar", "Foo de la O blah"},
	"Foo Bar Baz":          []string{"Fo Ba Ba", "Foo Bar Baz"},
	"Foo Bar Bar":          []string{"Foo Bar Bar", "Foo Bar Baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaar"},
	"bpauls":               []string{"bpauls", "2012-01-13"},
	"Peter Bar III":        []string{"Peter", "Peter Bar", "Peter Bar III"},
	"Peter III":            []string{"Peter", "Peter III", "Peter III Bar"},
}

func (s *PersonalSuite) TestGetBestFullName(c *C) {
	for best, names := range fullNameProvider {
		c.Assert(GetBestFullName(names), Equals, best)
	}
}
