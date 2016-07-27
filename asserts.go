package personal

import (
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
	"regexp"
)

func isFullNameCandidate(s string) bool {
	return !isEmail(s) && !isEmpty(s)
}

// Matches valid full names in a lax way
// (e.g. not considering capitalization).
var fullNameRegex, _ = regexp.Compile("^\\pL+?(\\s+(\\pL{1,2}\\.|\\pL+(-\\pL+)?))*(\\s[IV]+)?$")

func isWellFormedFullName(s string) bool {
	return fullNameRegex.MatchString(s)
}

var titlePrefixes = map[string]interface{}{
	"Mr":nil, "Ms":nil, "Mrs":nil, "Miss":nil, "Sir":nil, "Lady":nil, "Lord":nil,
	"Dr":nil, "Prof":nil,
}

func isTitlePrefix(s string) bool {
	_, present := titlePrefixes[s]
	return present
}

func containsNumbers(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}

	return false
}

func numberOfWords(s string) int {
	return len(strings.Split(s, " "))
}

func isEmail(s string) bool {
	return govalidator.IsEmail(s)
}

func isLowerCase(s string) bool {
	return govalidator.IsLowerCase(s)
}

func isEmpty(s string) bool {
	return len(s) == 0
}

var capitalizedFullNameRegex, _ =
	regexp.Compile("^\\p{Lu}\\p{Ll}*((\\s+\\p{Ll}{1,3}){0,2}\\s+(\\p{Lu}{1,2}\\.|\\p{Lu}\\p{Ll}*(-\\p{Lu}\\p{Ll}*)?))*(\\s[IV]+)?$")

func isCapitalizedFullName(s string) bool {
	return capitalizedFullNameRegex.MatchString(s)
}

func upcaseInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}

	return ""
}

func removeDots(str string) string {
	return strings.Replace(str, ".", "", -1)
}

func splitWords(str string) []string {
	return strings.Split(str, " ")
}

var preferredDomains = map[string]float64{
	"gmail.com":                2,
	"me.com":                   1,
	"live.com":                 1,
	"outlook.com":              1,
	"mail.ru":                  1,
	"qq.com":                   1,
	"hotmail.com":              -1,
	"users.noreply.github.com": -1,
}

func getPreferredDomainScore(email string) float64 {
	domain := getDomainFromEmail(email)
	score, preferred := preferredDomains[domain]
	if !preferred {
		return 0
	}
	return score
}

func isPrimaryDomain(email string) bool {
	domain := getDomainFromEmail(email)
	base := removeTLD(domain)

	parts := strings.Split(base, ".")
	if len(parts) == 1 {
		return true
	}

	return false
}

func removeTLD(domain string) string {
	parts := strings.Split(domain, ".")

	i := len(parts) - 1
	if i > 3 {
		i = 3
	}

	for {
		tld := strings.Join(parts[len(parts)-i:], ".")
		if _, valid := TLDs[tld]; valid || i < 1 {
			break
		}

		i--
	}

	return strings.Join(parts[:len(parts)-i], ".")
}

func getDomainFromEmail(email string) string {
	p := strings.Split(email, "@")
	if len(p) != 2 {
		return ""
	}

	return p[1]
}

func isValidTLD(email string) bool {
	tld := getTLDFromEmail(email)
	_, valid := TLDs[tld]

	return valid
}

func getTLDFromEmail(email string) string {
	p := strings.Split(email, ".")
	if len(p) < 2 {
		return ""
	}

	return p[len(p)-1]
}

func clean(s string) string {
	return govalidator.Trim(s, "")
}
