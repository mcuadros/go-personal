package personal

//GetBestFirstname returns the first name that can be found on the list of
//fullnames.
func GetBestFirstname(names []string) string {
	fn := GetBestFullName(names)
	fn = removeDots(fn)

	words := splitWords(fn)
	if len(words) == 1 {
		return ""
	}

	for idx, name := range words {
		if idx == 0 && isTitlePrefix(name) {
			continue
		}
		if len(name) > 1 {
			return name
		}
	}

	return ""
}
