package calculator

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectedErr bool
	}{
		{"valid division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative division", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.expectedErr && err == nil {
				t.Errorf("expected an error but got nil")
			}
			if !tt.expectedErr && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, result)
			}
		})
	}
}
