package personal

import (
	. "gopkg.in/check.v1"
)

var emailProvider = map[string][]string{
	"foo@bar.com": []string{"foo@bar.com", "foo"},

	//We prefer some domains againsts others
	"foo@gmail.com": []string{"foo@bar.com", "foo@gmail.com"},

	//Local domains
	"foo@baz.com": []string{"foo@bar.(none)", "foo@baz.com"},
	"foo@qux.com": []string{"foo@bar.local", "foo@qux.com"},
	"qux@qux.com": []string{"foo@bar.localdomain", "qux@qux.com"},

	//We prefer top domains againts subdomains
	"foo@foo.co.at.pn": []string{"foo@foo.bar.com", "foo@foo.co.at.pn"},
	"foo@bar.co.uk":    []string{"foo@bar.co.uk", "foo@foo.fooo.com"},
	"foo@bar.com.es":   []string{"foo@bar.com.es", "foo@bar.foo.fo.com"},
}

func (s *PersonalSuite) TestGetBestEmail(c *C) {
	for best, names := range emailProvider {
		c.Assert(best, Equals, GetBestEmail(names))
	}
}
