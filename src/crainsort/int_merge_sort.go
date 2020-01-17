package crainsort

func IntMergesort(arr []int) []int {
	size := len(arr)

	if size == 1 {
		return arr
	}

	mid := size / 2
	left := IntMergesort(arr[0:mid])
	right := IntMergesort(arr[mid:size])
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	merged := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for i < len(left) {
		merged = append(merged, left[i])
		i++
	}

	for j < len(right) {
		merged = append(merged, right[j])
		j++
	}

	return merged
}
