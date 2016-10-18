package personal

//GetBestFullName returns the best fullname that can be found on the list of
//strings. It gives priority to strings with correctly capitalized names, only
//letter strings, longer strings and strings with more than two words and less
//than four.
func GetBestFullName(names []string) string {
	return best(ScoreFullNames(names))
}

//ScoreFullNames returns a map with the fullnames and his score, higher is better
func ScoreFullNames(names []string) map[string]float64 {
	return score(names, ScoreFullName)
}

//ScoreFullName returns the full name score for a given string
func ScoreFullName(s string) float64 {
	if !isFullNameCandidate(s) || containsNumbers(s) {
		return -1
	}
	// Add up to 1 point for names up to 4 words. This covers most common
	// naming conventions, with few exception.
	//
	// In our use case, we found ~98.8% users use at most 3 names, and
	// 99.7% use at most 4.
	//
	// At 5 or more words, non full name and weird results would start
	// being the norm rather than the exception.
	var score float64
	if n := numberOfWords(s); n > 1 && n <= 4 {
		score += float64(n) / 4
	}

	if isLowerCase(s) {
		score -= .1
	}
	if isWellFormedFullName(s) {
		score++
	}
	if isCapitalizedFullName(s) {
		score++
	}

	// Prefer longer names if they are under a sane length limit.
	// ~99.9% names we found were, at most, 35 long.
	if l := len(s); l <= 35 {
		score += float64(l) / 35
	}

	return score
}

type scoreFunc func(string) float64

func score(ls []string, scorer scoreFunc) map[string]float64 {
	scores := map[string]float64{}
	for _, s := range ls {
		scores[s] = scorer(clean(s))
	}
	return scores
}

func best(scores map[string]float64) string {
	var higher float64
	var best string
	for name, score := range scores {
		if score >= higher {
			best, higher = name, score
		}
	}
	return best
}
