package main

func reverse(input []int64) (result []int64) {
	length := len(input)
	for i := length - 1; i >= 0; i-- {
		result = append(result, input[i])
	}
	return
}
