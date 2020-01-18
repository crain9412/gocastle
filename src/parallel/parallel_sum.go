package parallel

func Sum(arr []int) int {
	parentChannel := make(chan int)
	go Divide(arr, parentChannel)
	sum := <-parentChannel
	return sum
}

func Divide(arr []int, channel chan int) {
	if len(arr) <= 2 {
		channel <- Conquer(arr)
	} else {
		newChannel := make(chan int)
		newChannel2 := make(chan int)
		go Divide(arr[:len(arr)/2], newChannel)
		go Divide(arr[len(arr)/2:], newChannel2)
		newChannels := []chan int{newChannel, newChannel2}
		combinedSum := Combine(newChannels)
		channel <- combinedSum
	}
}

func Combine(inputs []chan int) int {
	sum := 0

	for i := 0; i < len(inputs); i++ {
		sum += <-inputs[i]
	}

	return sum
}

func Conquer(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
