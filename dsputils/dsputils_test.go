package dsputils

import (
	"testing"
)

type segmentTest struct {
	segs     int
	noverlap float64
	slices   [][]int
}

var segmentTests = []segmentTest{
	{
		3,
		.5,
		[][]int{
			{0, 8},
			{4, 12},
			{8, 16},
		},
	},
}

func TestSegment(t *testing.T) {
	x := make([]complex128, 16)
	for n := range x {
		x[n] = complex(float64(n), 0)
	}

	for _, st := range segmentTests {
		v := Segment(x, st.segs, st.noverlap)
		s := make([][]complex128, st.segs)
		for i, sl := range st.slices {
			s[i] = x[sl[0]:sl[1]]
		}

		if !PrettyClose2(v, s) {
			t.Error("Segment error: expected:", s, ", output:", v)
		}
	}
}
