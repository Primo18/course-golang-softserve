package calculator

import "testing"

func FuzzDivide(f *testing.F) {
	f.Add(10.0, 2.0)
	f.Add(1.0, 0.0)

	f.Fuzz(func(t *testing.T, a, b float64) {
		_, err := Divide(a, b)
		if b == 0 && err == nil {
			t.Errorf("expected an error for division by zero but got nil")
		}
	})
}
