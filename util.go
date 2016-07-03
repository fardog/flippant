package flippant

// ByLength exposes an interface for sorting strings by length
type ByLength []string

func (s ByLength) Len() int      { return len(s) }
func (s ByLength) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByLength) Less(i, j int) bool {
	return runeLen(s[i]) < runeLen(s[j])
}

// MakeLengthMap creates a map of [length]transitionPoints, given an array
// of strings sorted by length, shortest to longest
func MakeLengthMap(ss []string) map[int]int {
	m := make(map[int]int)

	for i, s := range ss {
		sl := runeLen(s)

		if i == 0 {
			m[sl] = i
		} else if sl > runeLen(ss[i-1]) {
			m[sl] = i
		}
	}

	return m
}

func runeLen(s string) int {
	return len([]rune(s))
}
