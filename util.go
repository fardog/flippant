package flippant

type byLength []string

func (s byLength) Len() int      { return len(s) }
func (s byLength) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byLength) Less(i, j int) bool {
	return runeLen(s[i]) < runeLen(s[j])
}

// MakeLengthMap creates a map of [length]transitionPoints, given an array
// of strings sorted by length, shortest to longest
func makeLengthMap(ss []string) map[int]int {
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
