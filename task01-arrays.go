package main

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	length := len(input)
	for _, i := range input {
		sum += i
	}
	return sum / float32(length)
}
