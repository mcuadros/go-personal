package personal

import (
	. "gopkg.in/check.v1"
)

var emailProvider = map[string][]string{
	// Invalid emails are ignored
	"foo@bar.com": {"foo@bar.com", "foo"},
	"":            {"foo"},

	//We prefer some domains againsts others
	"foo@gmail.com":  {"foo@bar.com", "foo@gmail.com"},
	"foo2@gmail.com": {"foo@live.com", "foo2@gmail.com"},
	"foo@live.com":   {"foo@bar.com", "foo@live.com"},

	"foo@hotmail.com":              {"foo@hotmail.com"},
	"foo@users.noreply.github.com": {"foo@users.noreply.github.com"},
	"foo@mydomain.com":             {"foo@mydomain.com", "foo@hotmail.com"},

	//Local domains
	"foo@baz.com": {"foo@bar.(none)", "foo@baz.com"},
	"foo@qux.com": {"foo@bar.local", "foo@qux.com"},
	"qux@qux.com": {"foo@bar.localdomain", "qux@qux.com"},

	//We prefer top domains against subdomains
	"foo@foo.co.at.pn": {"foo@foo.bar.com", "foo@foo.co.at.pn"},
	"foo@bar.co.uk":    {"foo@bar.co.uk", "foo@foo.fooo.com"},
	"foo@bar.com.es":   {"foo@bar.com.es", "foo@bar.foo.fo.com"},
}

func (s *PersonalSuite) TestGetBestEmail(c *C) {
	for best, emails := range emailProvider {
		c.Assert(GetBestEmail(emails), Equals, best)
	}
}
