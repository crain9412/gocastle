package parallel

import (
	"testing"
)

func Test_Parallel_Sum(t *testing.T) {
	arr := []int{7, 2, 8, -9, 4, 0}
	want := 12

	sum := Sum(arr)

	if got := sum; got != want {
		t.Errorf("Parallel sum was %q, want %q", got, want)
	}
}

func Benchmark_Parallel_Sum(b *testing.B) {
	arr := []int{0}

	for i := 0; i < b.N; i++ {
		arr = append(arr, i)
	}

	Sum(arr)
}

func Benchmark_Sequential_Sum(b *testing.B) {
	sum := 0
	arr := []int{0}

	for i := 0; i < b.N; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		sum += arr[i]
	}
}
