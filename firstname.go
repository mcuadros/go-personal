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

	for _, name := range words {
		if len(name) > 1 {
			return name
		}
	}

	return ""
}
