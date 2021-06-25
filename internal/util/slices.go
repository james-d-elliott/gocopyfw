package util

// IsStringInSlice checks if a haystack of strings has a needle.
func IsStringInSlice(needle string, haystack []string) (is bool) {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}

	return false
}
