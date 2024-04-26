package arrutil

// Contains checks if a string is present in a slice
func Contains(s []any, str any) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
