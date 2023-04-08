// Package hilbert can be used to obtain the node points on a Hilbert
// curve.
package hilbert

import "fmt"

// StepXY returns the XY coordinate of step, s, in an w by w node
// square (w = 2^(n+1)). This function will return an error if s >= w.
func StepXY(n, s uint) (x, y uint, err error) {
	w := uint(1) << (n + 1)
	if s >= w*w {
		err = fmt.Errorf("step %d is outside square %d*%d", s, w, w)
		return
	}
	for n, o := s, uint(1); o < w; o <<= 1 {
		xx := 1 & (n >> 1)
		yy := 1 & (n ^ xx)
		if yy == 0 {
			if xx == 1 {
				x = o - (1 + x)
				y = o - (1 + y)
			}
			x, y = y, x
		}
		x += o * xx
		y += o * yy
		n >>= 2
	}
	return
}
