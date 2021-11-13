package benchmarks

import (
	"testing"
)

// go test -v

func TestRecFib(t *testing.T) {
	data := [] struct {
		n uint
		want uint
	} {
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8},
		{10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := RecFib(d.n); got != d.want {
			t.Errorf("got: %d, want: %d", got, d.want)
		}
	}
}

func TestSeqFib(t *testing.T) {
	data := [] struct {
		n uint
		want uint
	} {
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8},
		{10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := SeqFib(d.n); got != d.want {
			t.Errorf("got: %d, want: %d", got, d.want)
		}
	}
}

func BenchmarkRecFib10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        RecFib(10)
    }
}


// go test -bench=RecFib20
func BenchmarkRecFib20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        RecFib(20)
    }
}