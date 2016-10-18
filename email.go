package personal

//GetBestEmail returns the best email that can be found on the list of
//strings.
func GetBestEmail(emails []string) string {
	return best(ScoreEmails(emails))
}

//ScoreEmails returns a map with the emails and its score, higher is better
func ScoreEmails(emails []string) map[string]float64 {
	return score(emails, ScoreEmail)
}

//ScoreEmail returns the score for a given email, higher is better.
func ScoreEmail(email string) float64 {
	if !isEmail(email) || !isValidTLD(email) {
		return -1
	}
	var score float64
	if isPrimaryDomain(email) {
		score++
	}
	return score + getPreferredDomainScore(email)
}
