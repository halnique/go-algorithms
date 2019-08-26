package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	cases := []struct {
		a []int
		e []int
	}{
		{
			[]int{1, 3, 5, 2, 4},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{100, 12345, 9, -4, -10, 0, 20},
			[]int{-10, -4, 0, 9, 20, 100, 12345},
		},
	}

	for _, c := range cases {
		b := Bubble{}
		r := b.Sort(c.a)
		assert.Equal(t, c.e, r, "Bubble(%q) == %q, expected %q", c.a, r, c.e)
	}
}
