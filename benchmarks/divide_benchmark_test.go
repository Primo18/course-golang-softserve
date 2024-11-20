package benchmarks

import (
	"myapp/internal/calculator"
	"testing"
)

// Benchmark simple
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := calculator.Divide(10, 2)
		// Evitar optimizaciones del compilador
		if err != nil || result == 0 {
			b.Fatal("unexpected result")
		}
	}
}

// Benchmark con diferentes valores
func BenchmarkDivideVariousInputs(b *testing.B) {
	inputs := [][2]float64{
		{10, 2},
		{100, 5},
		{1000, 10},
		{1, 0.5},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := inputs[i%len(inputs)]
		result, err := calculator.Divide(input[0], input[1])
		if err != nil || result == 0 {
			b.Fatal("unexpected result")
		}
	}
}

// Benchmark para división por cero
func BenchmarkDivideByZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := calculator.Divide(10, 0)
		if err == nil {
			b.Fatal("expected error")
		}
	}
}
