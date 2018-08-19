package hamming

import "errors"

// Distance comment
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Some error")
	}
	index := 0
	hDistance := 0
	for _, val := range a {
		if string(val) != string(b[index]) {
			hDistance++
		}
		index++
	}
	return hDistance, nil
}
