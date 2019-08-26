package sort

// Bubble return []int
func Bubble(s []int) []int {
	l := len(s)

	for {
		f := false

		for i := 0; i < l-1; i++ {
			ci := i
			ni := i + 1

			a := s[ci]
			b := s[ni]

			if a > b {
				s[ni] = a
				s[ci] = b
				f = true
			}
		}

		if f == false {
			break
		}
	}

	return s
}
