package crainsort

func IntQuicksort(arr []int) {
	size := len(arr)

	if size == 1 {
		return
	}

	pivot := Partition(arr)
	if pivot > 0 {
		IntQuicksort(arr[:pivot])
	}
	if pivot+1 < size {
		IntQuicksort(arr[pivot+1:])
	}
}

func Partition(arr []int) int {
	size := len(arr)
	right := size - 1
	mid := size / 2
	pivot := arr[mid]
	store := 0

	Swap(arr, mid, right)

	for i := 0; i < size; i++ {
		if arr[i] < pivot {
			Swap(arr, store, i)
			store++
		}
	}

	Swap(arr, right, store)

	return store
}

func Swap(arr []int, i, j int) {
	if i != j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}
