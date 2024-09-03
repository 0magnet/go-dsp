package spectral

import (
	"testing"

	"github.com/0magnet/go-dsp/dsputils"
)

type segmentTest struct {
	size,
	noverlap int
	out [][]float64
}

var segmentTests = []segmentTest{
	{
		4, 0,
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	},
	{
		4, 1,
		[][]float64{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
			{7, 8, 9, 10},
		},
	},
	{
		4, 2,
		[][]float64{
			{1, 2, 3, 4},
			{3, 4, 5, 6},
			{5, 6, 7, 8},
			{7, 8, 9, 10},
		},
	},
}

func TestSegment(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range segmentTests {
		o := Segment(x, v.size, v.noverlap)
		if !dsputils.PrettyClose2F(o, v.out) {
			t.Error("Segment error\n  output:", o, "\nexpected:", v.out)
		}
	}
}
