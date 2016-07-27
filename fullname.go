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
	if !isFullNameCandidate(s) {
		score = -1
		return
	}

	if containsNumbers(s) {
		score = -1
		return
	}

	// Add up to 1 point for names up to 4 words. This covers most common
	// naming conventions, with few exception.
	//
	// In our use case, we found ~98.8% users use at most 3 names, and
	// 99.7% use at most 4.
	//
	// At 5 or more words, non full name and weird results would start
	// being the norm rather than the exception.
	nWords := numberOfWords(s)
	if nWords > 1 && nWords <= 4 {
		score += float64(nWords) / 4.
	}

	if isLowerCase(s) {
		score -= .1
	}

	if isWellFormedFullName(s) {
		score += 1
	}

	if isCapitalizedFullName(s) {
		score += 1
	}

	// Prefer longer names if they are under a sane length limit.
	// ~99.9% names we found were, at most, 35 long.
	length := len(s)
	if length <= 35 {
		score += float64(length) / 35
	}

	return
}
