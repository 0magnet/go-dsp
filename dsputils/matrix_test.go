package dsputils

import (
	"testing"
)

func TestMakeMatrix(t *testing.T) {
	m := MakeMatrix(
		[]complex128{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 0, 1, 2,

			3, 4, 5, 6,
			7, 8, 9, 0,
			4, 3, 2, 1},
		[]int{2, 3, 4})

	checkArr(t, m.Dim([]int{1, 0, -1}), ToComplex([]float64{3, 4, 5, 6}))
	checkArr(t, m.Dim([]int{0, -1, 2}), ToComplex([]float64{3, 7, 1}))
	checkArr(t, m.Dim([]int{-1, 1, 3}), ToComplex([]float64{8, 0}))

	s := ToComplex([]float64{10, 11, 12})
	i := []int{1, -1, 3}
	m.SetDim(s, i)
	checkArr(t, m.Dim(i), s)

	v := complex(14, 0)
	m.SetValue(v, i)
	checkFloat(t, m.Value(i), v)
}

func checkArr(t *testing.T, have, want []complex128) {
	if !PrettyCloseC(have, want) {
		t.Error("have:", have, "want:", want)
	}
}

func checkFloat(t *testing.T, have, want complex128) {
	if !ComplexEqual(have, want) {
		t.Error("have:", have, "want:", want)
	}
}
