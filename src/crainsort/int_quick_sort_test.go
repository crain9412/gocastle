package crainsort

import "testing"

func Test_IntQuickSort(t *testing.T) {
	starting := []int{4, 1, 0, 2, 3}
	IntQuicksort(starting)
	for i := 0; i < len(starting); i++ {
		if i != starting[i] {
			t.Errorf("Sorted element was %q, want %q", starting[i], i)
		}
	}
}
