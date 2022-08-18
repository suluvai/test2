package strings

import "unicode"

// ToUpper uppercases all the runes in its argument string.
func ToUpperString(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}
