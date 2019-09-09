// Package hamming calculate Hamming distance between DNA strands.
// DNA strands are strings composed of A, G, C, T.
package hamming

import (
	"github.com/pkg/errors"
)

// Distance calculate Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("cannot compute hamming distance on different length sequences")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
