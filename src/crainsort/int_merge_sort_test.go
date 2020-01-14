package crainsort

import "testing"

func Test_IntMergeSort(t *testing.T) {
	starting := []int{4, 1, 0, 2, 3}
	sorted := IntMergesort(starting)
	for i := 0; i < len(sorted); i++ {
	    if (i != sorted[i]) {
	        t.Errorf("Sorted element was %q, want %q", sorted[i], i)
	    }
	}
}