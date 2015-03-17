//A very small library to score strings as fullname. Given a list of strings the
//library  returns which is the most accurate Fullname.
package fullname

import (
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
)

//GetBestFullName returns the best fullname that can be found on the list of
//strings. Its gives priority to strings with correct capitalize names, only
//letters strings, longer strings and strings with more than two words and less
//than four.
func GetBestFullName(names []string) string {
	var higher float64
	var best string
	for name, score := range ScoreFullNames(names) {
		if score > higher {
			higher = score
			best = name
		}
	}

	return best
}

//ScoreFullNames returns a map with the strings and his score, higher is better
func ScoreFullNames(names []string) map[string]float64 {
	r := make(map[string]float64, 0)
	for _, name := range names {
		name = clean(name)
		r[name] = scoreFullname(name)
	}

	return r
}

func scoreFullname(s string) (score float64) {
	if !isFullname(s) {
		score = -1
		return
	}

	if !isSingleWord(s) && hasLessWordsThan(s, 4) {
		score += 1
	}

	if isLowerCase(s) {
		score -= .1
	}

	if isOnlyLetters(s) {
		score += 1
	}

	if isUpCaseInitial(s) {
		score += 1
	}

	score += float64(len(s)) / 100

	return
}

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

func clean(s string) string {
	return govalidator.Trim(s, "")
}
