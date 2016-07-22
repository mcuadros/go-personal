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
		r[email] = ScoreEmail(email)
	}

	return r
}

//ScoreEmail returns the score for a given email, higher is better.
func ScoreEmail(email string) (score float64) {
	if !isEmail(email) || !isValidTLD(email) {
		score = -1
		return
	}

	if isPrimaryDomain(email) {
		score += 1
	}

	score += getPreferredDomainScore(email)

	return
}
