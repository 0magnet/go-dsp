// Package spectral provides spectral analysis functions for digital signal processing.
package spectral

// Segment x segmented into segments of length size with specified noverlap.
// Number of segments returned is (len(x) - size) / (size - noverlap) + 1.
func Segment(x []float64, size, noverlap int) [][]float64 {
	stride := size - noverlap
	lx := len(x)

	var segments int
	if lx == size {
		segments = 1
	} else if lx > size {
		segments = (len(x)-size)/stride + 1
	} else {
		segments = 0
	}

	r := make([][]float64, segments)
	for i, offset := 0, 0; i < segments; i++ {
		r[i] = make([]float64, size)

		for j := 0; j < size; j++ {
			r[i][j] = x[offset+j]
		}

		offset += stride
	}

	return r
}
