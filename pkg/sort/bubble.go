package sort

type Bubble struct {}

// Sort return []int
func (*Bubble) Sort(s []int) []int {
	r := append([]int{}, s...)
	l := len(s)

	for {
		f := false

		for i := 0; i < l - 1; i++ {
			ci := i
			ni := i + 1

			ia := r[ci]
			ib := r[ni]

			if ia > ib {
				r[ni] = ia
				r[ci] = ib
				f = true
			}
		}

		if f == false {
			break
		}
	}

	return r
}
