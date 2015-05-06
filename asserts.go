package personal

import (
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
)

func isFullname(s string) bool {
	return !isEmail(s) && !isEmpty(s)
}

func isOnlyLetters(s string) bool {
	woSpaces := strings.Replace(s, " ", "", -1)
	return govalidator.IsUTFLetter(woSpaces)
}

func isSingleWord(s string) bool {
	return len(strings.Split(s, " ")) == 1
}

func hasLessWordsThan(s string, n int) bool {
	return len(strings.Split(s, " ")) < n
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

func isUpCaseInitial(s string) bool {
	return s == upcaseInitialAllWords(strings.ToLower(s))
}

func upcaseInitialAllWords(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		words[i] = upcaseInitial(w)
	}

	return strings.Join(words, " ")
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

var preferredDomains = map[string]interface{}{"gmail.com": nil}

func isPreferredDomain(email string) bool {
	domain := getDomainFromEmail(email)
	_, preferred := preferredDomains[domain]

	return preferred
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
