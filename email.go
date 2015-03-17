//A very small library to score strings as fullname. Given a list of strings the
//library  returns which is the most accurate Fullname.
package personal

//GetBestFullName returns the best email that can be found on the list of
//strings.
func GetBestEmail(names []string) string {
	var higher float64
	var best string

	for name, score := range ScoreEmails(names) {
		if score >= higher {
			higher = score
			best = name
		}
	}

	return best
}

//ScoreEmails returns a map with the emails and his score, higher is better
func ScoreEmails(emails []string) map[string]float64 {

	r := make(map[string]float64, 0)
	for _, email := range emails {
		email = clean(email)
		r[email] = scoreEmail(email)
	}

	return r
}

func scoreEmail(s string) (score float64) {
	if !isEmail(s) || !isValidTLD(s) {
		score = -1
		return
	}

	if isPrimaryDomain(s) {
		score += 1
	}

	if isPreferredDomain(s) {
		score += 1
	}

	return
}
