package split

import "strings"

// Split ...
func Split(s string, sep string) []string {
	out := make([]string, 0, len([]rune(s)))
	index := strings.Index(s, sep)
	for index >= 0 {
		out = append(out, s[:index])
		s = s[index+1:]
		index = strings.Index(s, sep)
	}
	out = append(out, s)
	return out
}
