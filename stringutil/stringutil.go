package stringutil

import "regexp"

func MatchScidPattern(input string) (bool, error) {
	match, err := regexp.MatchString(`^SL\d{4}$`, input)
	return match, err
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
