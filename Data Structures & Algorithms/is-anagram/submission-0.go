import "slices"
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes, tRunes := []rune(s), []rune(t)
	slices.Sort(sRunes)
	slices.Sort(tRunes)
	return string(sRunes) == string(tRunes)
}
