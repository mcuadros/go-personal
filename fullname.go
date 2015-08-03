package personal

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

//ScoreFullNames returns a map with the fullnames and his score, higher is better
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

	if isOnlyValidChars(s) {
		score += 1
	}

	if isUpCaseInitial(s) {
		score += 1
	}

	score += float64(len(s)) / 100

	return
}
