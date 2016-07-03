package flippant

import "testing"

func TestMakeLengthMap(t *testing.T) {
	ss := []string{"o", "i", "oo", "iii", "vvv", "vvvv", "vvvv"}
	expected := map[int]int{
		1: 0,
		2: 2,
		3: 3,
		4: 5,
	}

	lm := MakeLengthMap(ss)

	for k, v := range lm {
		if expected[k] != v {
			t.Errorf("unexpected value %d for key %d", v, k)
		}
	}
}
