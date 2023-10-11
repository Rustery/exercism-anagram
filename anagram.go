package anagram

import "strings"

func isAnagram(subject string, candidate string) bool {

	if subject == candidate || len(subject) != len(candidate) {
		return false
	}

	for _, v := range subject {
		if strings.Count(subject, string(v)) != strings.Count(candidate, string(v)) {
			return false
		}
	}

	return true
}

func Detect(subject string, candidates []string) []string {
	var results []string
	for _, v := range candidates {
		if isAnagram(strings.ToLower(subject), strings.ToLower(v)) {
			results = append(results, v)
		}
	}
	return results
}
