package ch34

import "testing"

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {

		ret := square(inputs[i])

		if ret != expected[i] {
			t.Errorf("input is %d", inputs[i])
		}

	}
}
